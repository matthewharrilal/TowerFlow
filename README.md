# Tower Flow ⚡️
Tower flow enables users to be able to message individual users **concurrently** 📲. Built upon Twillio SMS API and utilizing [GORM](https://github.com/jinzhu/gorm) (Go Object Relational Mapper) to persist data, this was and is an example of understanding concurrency in a real world context!

### Prerequisites
In order to **interact** with the Twillio API to send messages you will need proper authentication credentials 👮, **Account SID** and and **Authentication Token**.

Navigate over to the [Twillio Console](https://www.twilio.com/console) and create an application with SMS messaging 📲!

Once doing so you should now have been issued an **authentication token** and **account sid** number. It should look a bit like this 🔥!

![Account Credentials](https://github.com/matthewharrilal/Concurrent-SMS-Messaging/blob/master/Assets/Twillio-Console.png)

**HIGHLY suggest** once you have obtained these credentials to place them in a **.env** file and stored in a **git ignore**. 

Once added SMS capabilities to your application on Twillio you should have received a telephone number that will act as the source of all outgoing messages! If you did not receive one navigate to phone numbers [section](https://www.twilio.com/console/phone-numbers/incoming) of the Twillio console.

### Installing
To install Tower Flow execute this command
 **_go get https://github.com/matthewharrilal/Concurrent-SMS-Messaging_**

 ##### What does a Message Object structure look like?
```
type Message struct {
	gorm.Model // Embeds schema in db with auto incrementing ID, created at, updated at, and deleted at attributes

	DateCreated string `json:"date_created"`

	MessageDirection string `json:"direction"` // Whether we are receiving or sending

	AccountIdentifier string `json:"account_sid"` 

	MessageIdentifier string `json:"sid"`

	Body string `json:"body"`

	NumberOfSegments string `json:"num_segments"` // Number of components within message
}
```

```

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
	clientManager.SendMessages(destinationNumbers, messageContent, messageChannel)
}

```

## Built With

* [Dropwizard](http://www.dropwizard.io/1.0.2/docs/) - The web framework used
* [Maven](https://maven.apache.org/) - Dependency Management
* [ROME](https://rometools.github.io/rome/) - Used to generate RSS Feeds

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Billie Thompson** - *Initial work* - [PurpleBooth](https://github.com/PurpleBooth)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc

