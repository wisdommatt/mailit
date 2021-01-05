# Mailit
An email kit for golang to be reused across different projects.

## Introduction
Mailit is an easy to use mail library to send emails.

## Features
Mailit supports:
- Plain text emails
- Text / HTML template based emails
- Attachments
- Sending emails to multiple recipients

## Requirements
> * Go 1.5

## Documentation

https://pkg.go.dev/github.com/wisdommatt/mailit
***
We also recommend you use the how to guide on this page because Mailit does basically two things: 

* Send Plain Text Emails
* Send Template Based Emails

## How To Use Mailit
######  Feel free to copy and paste the codes

***
- #### Step 1: Initialize a mailer variable and pass in SMTP configs.
    > #### <span style="color:green"> Note: This is usually done just once<span>

    ```go
        smtpConf := mailit.SMTPConfig{
            Host:     "domain.com",
            Port:     563,
            Username: "username@domain.com",
            Password: "*********",
        }
        mailer := mailit.NewMailer(smtpConf)
    ```

- #### Step 2: Send email 
    > ##### Sending plain text email

    ```go
        textDep := mailit.TextDependencies{
            From:        "sender@domain.com",
            To:          []string{"user1@domain.com", "user2@domain.com", "user3@domain.com"},
            Subject:     "Welcome to Mailit",
            Body:        "Message Body",
            Attachments: []string{"attachments/1.txt", "attachments/2.txt"},
        }
        err := mailer.SendText(textDep)
    ```
    ***

    > ##### Sending HTML template email

    ```go
        tempDep := mailit.TemplateDependencies{
            From:        "sender@domain.com",
            To:          []string{"user1@domain.com", "user2@domain.com", "user3@domain.com"},
            Subject:     "Welcome to Mailit",
            ContentType: "text/html",
            Template: "templates/welcome.html",
            TemplateData: "Any data",
            Attachments: []string{"attachments/1.txt", "attachments/2.txt"},
        }
        err := mailer.SendTemplate(tempDep)
    ```

    ***

    > ##### Sending Text template email

    ```go
        tempDep := mailit.TemplateDependencies{
            From:        "sender@domain.com",
            To:          []string{"user1@domain.com", "user2@domain.com", "user3@domain.com"},
            Subject:     "Welcome to Mailit",
            ContentType: "text/plain",
            Template: "templates/sample.txt",
            TemplateData: struct{
                Name, Email string
            }{
                Name: "Wisdom Matt",
                Email: "user@example.com",
            },
            Attachments: []string{"attachments/1.txt", "attachments/2.txt"},
        }
        err := mailer.SendTemplate(tempDep)
    ```
***
> NOTE: You will access the template data the way it is access usually in Go templates e.g {{ .Name }} or {{ .Email }}
***
## Dependencies

- https://github.com/go-gomail/gomail

## Testing

This package uses manual testing because the gomail package which it depends on used core types instead of interfaces which in turn makes it difficult to test, a thorough test is done before any changed is made on the project so there is nothing to be scared of.

## License

[MIT](LICENSE)

## Author / Contact

[Wisdom Matthew](https://wisdommatt.github.io/) - github.com/wisdommatt

Feel free to email me at talk2wisdommatt@gmail.com