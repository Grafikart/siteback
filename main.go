package main

import (
	"compress/gzip"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"time"
)

func main() {
	err := godotenv.Load(".env.local")
	crash("erreur lors du chargement du fichier .env.local, %v", err)
	stdout, err := dump()
	crash("erreur lors de l'export de la base", err)
	stdout, err = compress(stdout)
	crash("erreur lors de la coompression", err)
	path, err := upload(stdout, os.Args[1])
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

// Compresse un io.Reader
func compress(input io.Reader) (io.Reader, error) {
	r, w := io.Pipe()
	go func() {
		defer w.Close()
		zip, _ := gzip.NewWriterLevel(w, gzip.BestCompression)
		zip.Name = filename() + ".sql"
		defer zip.Close()
		io.Copy(zip, input)
	}()
	return r, nil
}

// Upload sur S3
func upload(f io.Reader, bucket string) (string, error) {
	cfg := aws.NewConfig().
		WithRegion("fr-par").
		WithEndpoint("https://s3.fr-par.scw.cloud").
		WithCredentials(credentials.NewStaticCredentials(os.Getenv("S3_KEY"), os.Getenv("S3_SECRET"), ""))
	sess := session.Must(session.NewSession(cfg))
	uploader := s3manager.NewUploader(sess)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename() + ".gz"),
		Body:   f,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}
	return result.Location, nil
}

// Génère un dump SQL à partir du chemin fournit par symfony
func dump() (io.Reader, error) {
	regex, _ := regexp.Compile(`mysql://([^:]+):([^@]+)[^/]+/(\w*)`)
	matches := regex.FindStringSubmatch(os.Getenv("DATABASE_URL"))
	cmd := exec.Command("mysqldump", "-u"+matches[1], "-p"+matches[2], matches[3])
	r, w := io.Pipe()
	go func() {
		cmd.Stdout = w
		defer w.Close()
		cmd.Run()
	}()
	return r, nil
}
