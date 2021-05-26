## API reusable SMTP Email Example

{
    "from" : "notification.to.admin@labirin-hybrid.com",
    "to" : "anita.mailist@gmail.com",
    "subject" : "Roy Menuliskan Pesan di Undangan Anda",
    "message" : "Dear Wahyu Ade Pratama,<br> <br>Just to explain things <b>easily</b>, In the above snippet, we have written all the smtp and email credentials in the main function, Though in a production app you should always use env variables for configurations. You can check Viper to manage configurations in production apps.",
    "smtp" : "labirin-hybrid.com",
    "port" : 465,
    "email" : "notification.to.admin@labirin-hybrid.com",
    "password" : "write your password here"
}
