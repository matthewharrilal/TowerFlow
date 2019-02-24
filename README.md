# Tower Flow ⚡️
Tower flow enables users to be able to message individual users **concurrently** 📲. Built upon Twillio SMS API and utilizing [GORM](https://github.com/jinzhu/gorm) (Go Object Relational Mapper) to persist data, this was and is an example of understanding concurrency in a real world context! This project is in active development and completely open source! If you have questions or ways to make this project better lets collaborate 👋! Either through a **pull request** or an email at **matthewharrilal@gmail.com** 📫!

### Prerequisites
In order to **interact** with the Twillio API to send messages you will need proper authentication credentials 👮, **Account SID** and and **Authentication Token**.

Navigate over to the [Twillio Console](https://www.twilio.com/console) and create an application with SMS messaging capabilities 📲!

Once done doing so you should now have been issued an **authentication token** and **account sid** number. It should look a bit like this 🔥!

![Account Credentials](https://github.com/matthewharrilal/Concurrent-SMS-Messaging/blob/master/Assets/Twillio-Console.png)

**HIGHLY suggest** once you have obtained these credentials to place them in a **.env** file and stored in a **git ignore** 
🤫. 

Once added SMS capabilities to your application on Twillio you should have received a telephone number ☎️ that will act as the source of all outgoing messages! If you did not receive one navigate to phone numbers [section](https://www.twilio.com/console/phone-numbers/incoming) of the Twillio console.

### Installing
To install Tower Flow ⚡️ execute this command
 **_go get https://github.com/matthewharrilal/Concurrent-SMS-Messaging_**

 ##### What does a Message Object structure look like?
``` go
type Message struct { 
	gorm.Model // Embeds schema in db with auto incrementing ID, created at, updated at, and deleted at attributes

	DateCreated string `json:"date_created"`

	MessageDirection string `json:"direction"` // Whether we are receiving or sending

	AccountIdentifier string `json:"account_sid"` 

	MessageIdentifier string `json:"sid"`

	Body string `json:"body"` // Unique identifier corresponding to the message object from Twillio

	NumberOfSegments string `json:"num_segments"` // Number of components within message
}
```
##### How do I send messages then?

``` go

func main() {

        // Storing Account Credentials in .env file
	err := godotenv.Load() // First load environment variables file
	if err != nil {
		log.Fatal(err)
	}

        // Your choice as to whether you want to persist messages or not!
	// ConfigureDatabase()


        // Supply the collection of telephone numbers that the outgoing message is going to be sent to!

	destinationNumbers := []string{"**********", "**********", "**********"}

        // Instantiate Channel that the formulated Message Objects are going to be sent through!
        messageChannel := make(chan Message) 

	// Pass in credentials that you were issued from the Twillio Console

	accountSID, authToken := os.Getenv("ACCOUNT_SID"), os.Getenv("AUTH_TOKEN")

	sourceNumber := os.Getenv("SOURCE_NUMBER") // Twillio Number that was issued

	// Your choice of client to execute the request used ... default is the http.DefaultClient

	clientManager := NewClient(nil, sourceNumber, authToken, accountSID)
    
        // Construct Message Contents and then call the Send Messages method!

        messageContent := "Any message you want!"
	message := clientManager.SendMessages(destinationNumbers, messageContent, messageChannel)


	// If you decided to persist messages you can then call the function
	PostMessage(&message, destinationNumbers)
}

```

## Built With

* [GORM](https://github.com/jinzhu/gorm) - Object Relational Mapper (With the use of an SQLite Database)
* [Twillio Services](https://www.twilio.com/) - API Contact and SMS Capabilities
* [GODOTENV](https://github.com/joho/godotenv) - Used to store secrets such as Twillio authentication credentials and source number



## Next Features
1. Able to send a collection of messages to collection of telephone numbers concurrently! Be able to map specific message to corresponding number!

2. Add intuitive API layer for the persistence of messages

3. Formulate more **robust testing**! 

## Authors

* **Matthew Harrilal** - *Looking For Software Engineering Internships!!* - [LinkedIn](https://www.linkedin.com/in/matthew-harrilal-b38377111/)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Special thanks to my backend instructor who has done an excellent job of teaching us golang and inspiration/application for this utility! Shoutout to [Droxey](https://github.com/droxey)

* Shoutout to [Subosito](https://github.com/subosito/twilio) for how to formulate the Client structure which contains relevant information toward the authenticity of the client!

