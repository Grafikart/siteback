package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	exporter "siteback/exporter"
	pipe "siteback/pipe"
	tool "siteback/tool"
	"time"
)

func main() {
	err := godotenv.Load(".env.local")
	crash("erreur lors du chargement du fichier .env.local, %v", err)

	stdout, err := tool.DumpDB()
	crash("erreur lors de l'export de la base", err)

	stdout, err = pipe.Gzip(stdout, filename()+".sql")
	crash("erreur lors de la coompression", err)

	path, err := exporter.S3(stdout, os.Args[1], filename()+".gz")
	// path, err := exporter.File(stdout, filename() + ".gz")
	crash("erreur lors de l'envoie sur S3", err)

	fmt.Printf("Sauvegarde effectuée : %q", path)
}

// Plante en cas d'erreur
func crash(message string, err error) {
	if err != nil {
		log.Fatalf(message+", %v", err)
	}
}

// Génère un nom de fichier lié à la date
func filename() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
}
