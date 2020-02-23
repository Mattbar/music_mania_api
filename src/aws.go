package main

import(
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	
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

// PutDB puts Item In the Music DB
func PutDB() {
	sess := session.Must(session.NewSession())
	creds := stscreds.NewCredentials(sess, "arn:aws:iam::080142785180:role/DDBFull")

	dsvc := dynamodb.New(sess, &aws.Config{Credentials: creds, Region: aws.String(S3Region)})

	input := &dynamodb.PutItemInput{
    Item: map[string]*dynamodb.AttributeValue{
        "Genre": {
            S: aws.String("genre1"),
        },
        "Artist": {
            S: aws.String("Fake Artist1"),
        },
        "Album": {
            S: aws.String("Artist1 Fake Album 1"),
		},
		"Song":{
			S: aws.String("Album1 Fake Song1"),
		},
		"S3Key":{
			S: aws.String("genre1/Fake Artist1/Artist1 Fake Album 1/Album1 Fake Song1"),
		},
    },
    ReturnConsumedCapacity: aws.String("TOTAL"),
    TableName:              aws.String("Music"),
	}

	result, err := dsvc.PutItem(input)
	
	if err != nil {
    if aerr, ok := err.(awserr.Error); ok {
    	switch aerr.Code() {
        case dynamodb.ErrCodeConditionalCheckFailedException:
            log.Println(dynamodb.ErrCodeConditionalCheckFailedException, aerr.Error())
        case dynamodb.ErrCodeProvisionedThroughputExceededException:
            log.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
        case dynamodb.ErrCodeResourceNotFoundException:
            log.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
        case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
            log.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
        case dynamodb.ErrCodeTransactionConflictException:
            log.Println(dynamodb.ErrCodeTransactionConflictException, aerr.Error())
        case dynamodb.ErrCodeRequestLimitExceeded:
            log.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
        case dynamodb.ErrCodeInternalServerError:
            log.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
        default:
            log.Println(aerr.Error())
        }
    } else {
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        log.Println(err.Error())
    }
    return
	}

	log.Println("Result",result)
}

// func GetDB(){
//     sess := session.Must(session.NewSession())
// 	creds := stscreds.NewCredentials(sess, "arn:aws:iam::080142785180:role/DDBFull")

//     dsvc := dynamodb.New(sess, &aws.Config{Credentials: creds, Region: aws.String(S3Region)})
    
//     input := &dynamodb.QueryInput{
//     ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
//         ":v1": {
//             S: aws.String("No One You Know"),
//         },
//     },
//     KeyConditionExpression: aws.String("Artist = :v1"),
//     ProjectionExpression:   aws.String("SongTitle"),
//     TableName:              aws.String("Music"),
//     }
// }




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
