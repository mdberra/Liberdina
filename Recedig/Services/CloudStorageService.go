package Services

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type CloudStorageService struct {
}

var (
	path string = "C:/GoogleCloudPlatform/"
	pathGCP      string = "./"

	projectID string = "LiberdinaT"
	bucket    string = "pepeya-bucket"
	publicKey string = "Keys/silken-vial-220421-9ea1970c49a9.json"
	pathFiles string = "Temp/"
	empresa   string = "Servina/"
	client    *storage.Client
	ctx       context.Context
	err       error
)

//func (css *CloudStorageService) Connect() (*storage.Client, error) {
//	return client, err
//}
func (css *CloudStorageService) GetClient() *storage.Client {
	return client
}
func (css *CloudStorageService) Connect() error {
	ctx = context.Background()
	client, err = storage.NewClient(ctx, option.WithCredentialsFile(path+publicKey))
	return err
}

func (css *CloudStorageService) UploadFile(fileName string, file multipart.File) error {
	wc := client.Bucket(bucket).Object(empresa + fileName).NewWriter(ctx)

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path+pathFiles+fileName, data, 777)
	if err != nil {
		return err
	}
	fdata, err := os.Open(path + pathFiles + fileName)
	if err != nil {
		return err
	}
	if _, err = io.Copy(wc, fdata); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	fdata.Close()
	if err := os.Remove(path + pathFiles + fileName); err != nil {
		return err
	}
	return nil
}

func (css *CloudStorageService) Read(fileName string) ([]byte, error) {
	rc, err := client.Bucket(bucket).Object(empresa + fileName).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(rc)
	defer rc.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (css *CloudStorageService) List() error {
	it := client.Bucket(bucket).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", attrs.Name)
	}
	return nil
}
func (css *CloudStorageService) ListByPrefix(w io.Writer, prefix, delim string) error {
	// Prefixes and delimiters can be used to emulate directory listings.
	// Prefixes can be used filter objects starting with prefix.
	// The delimiter argument can be used to restrict the results to only the
	// objects in the given "directory". Without the delimeter, the entire  tree
	// under the prefix is returned.
	//
	// For example, given these blobs:
	//   /a/1.txt
	//   /a/b/2.txt
	//
	// If you just specify prefix="a/", you'll get back:
	//   /a/1.txt
	//   /a/b/2.txt
	//
	// However, if you specify prefix="a/" and delim="/", you'll get back:
	//   /a/1.txt
	it := client.Bucket(bucket).Objects(ctx, &storage.Query{
		Prefix:    prefix,
		Delimiter: delim,
	})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Fprintln(w, attrs.Name)
	}
	// [END storage_list_files_with_prefix]
	return nil
}
func (css *CloudStorageService) Attrs(object string) (*storage.ObjectAttrs, error) {
	o := client.Bucket(bucket).Object(object)
	attrs, err := o.Attrs(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("Bucket: %v\n", attrs.Bucket)
	log.Printf("CacheControl: %v\n", attrs.CacheControl)
	log.Printf("ContentDisposition: %v\n", attrs.ContentDisposition)
	log.Printf("ContentEncoding: %v\n", attrs.ContentEncoding)
	log.Printf("ContentLanguage: %v\n", attrs.ContentLanguage)
	log.Printf("ContentType: %v\n", attrs.ContentType)
	log.Printf("Crc32c: %v\n", attrs.CRC32C)
	log.Printf("Generation: %v\n", attrs.Generation)
	log.Printf("KmsKeyName: %v\n", attrs.KMSKeyName)
	log.Printf("Md5Hash: %v\n", attrs.MD5)
	log.Printf("MediaLink: %v\n", attrs.MediaLink)
	log.Printf("Metageneration: %v\n", attrs.Metageneration)
	log.Printf("Name: %v\n", attrs.Name)
	log.Printf("Size: %v\n", attrs.Size)
	log.Printf("StorageClass: %v\n", attrs.StorageClass)
	log.Printf("TimeCreated: %v\n", attrs.Created)
	log.Printf("Updated: %v\n", attrs.Updated)
	log.Printf("Event-based hold enabled? %t\n", attrs.EventBasedHold)
	log.Printf("Temporary hold enabled? %t\n", attrs.TemporaryHold)
	log.Printf("Retention expiration time %v\n", attrs.RetentionExpirationTime)
	log.Print("\n\nMetadata\n")
	for key, value := range attrs.Metadata {
		log.Printf("\t%v = %v\n", key, value)
	}
	return attrs, nil
}
