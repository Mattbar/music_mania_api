package main

import(
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	
	"log"
	"time"
)
// s3 constants
const (
  S3Region = "us-east-1"
	S3Bucket = "mbarnesbucket1"
)

// GetURL gets the url for the song
func GetURL(key string) string{
	log.Println("GET URL")
	sess := session.Must(session.NewSession())
	creds := stscreds.NewCredentials(sess, "arn:aws:iam::080142785180:role/s3Full")
	
	svc := s3.New(sess, &aws.Config{Credentials: creds, Region: aws.String(S3Region)})
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
 		Bucket: aws.String(S3Bucket),
    Key:    aws.String(key),
  })
	
	urlStr, err := req.Presign(15 * time.Minute)

    if err != nil {
        log.Println("Failed to sign request", err)
    }

		log.Println("The URL is", urlStr)
	return urlStr
}




// -------------------upload code-----------------------------------------
	// "github.com/aws/aws-sdk-go/awserr"
	//"github.com/aws/aws-sdk-go/service/s3/s3manager"
	//"fmt"
// func addToBucket(paths []string) {

// 	  // Create a single AWS session
// 		sess := session.Must(session.NewSession())
// 		creds := stscreds.NewCredentials(sess, "arn:aws:iam::080142785180:role/s3Full")


	
		
// 		// paths := [3]string{"/home/mattbarnes/Music/FakeArtist1/FakeAlbum1/3788-funkorama-by-kevin-macleod.mp3","/home/mattbarnes/Music/FakeArtist1/FakeAlbum1/bensound-ukulele.mp3","/home/mattbarnes/Music/FakeArtist1/FakeAlbum2/Super Mario Bros. Theme Song.mp3"}

// 		for i := 0; i < len(paths); i++ {
// 			log.Println(paths[i])
			
// 			path := paths[i]
// 			iter := NewDirectoryIterator(S3Bucket, path)
// 			upload := s3manager.NewUploader(session.New(&aws.Config{Credentials: creds, Region: aws.String(S3Region)}))

// 			err := upload.UploadWithIterator(aws.BackgroundContext(), iter) 
		
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Printf("Successfully uploaded %q to %q", path, S3Bucket)

// 		}
// }
