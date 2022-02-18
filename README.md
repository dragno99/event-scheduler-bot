# Event scheduler bot for discord
A simple Discord bot created in golang that allows admin of a server to create scheduled google calender events. This bot will help admin to schedule an event with auto generated google's meet link and add attendees and feed this event to attendees's google calender.


### Functions of bot
 * Create and schedule an event.
 * Show upcoming events.
 * Add attendees in an already scheduled event.
 * Delete an event.

### Create and schedule an event
 * To create an event , input data will be in the following format.
 * First line starts with `!schedule` 
 * Second line will contain the summary of event.
 * Third line will contain the start date and time of event. `YYYY-MM-DDTHH-MM-SS` , i.e `2022-02-22T09-00-00`
 * Fourth line will contain the end date and time of event. `YYYY-MM-DDTHH-MM-SS` , i.e `2022-02-22T11-00-00`
 * Fifth line will contain the email address of attendees seperated by spaces. i.e `xyz@gmai.com abx@gmail.com pqr@work.ac.in`

### Show upcoming events.
 * To see upcoming events, enter command `!upcoming`

### Add attendees in an already scheduled event.
 * To add attendees in an already scheduled event , first you need to see upcoming event by the `!upcoming` command.
 * After getting the list of upcoming event, You can add any number of attendes by the following command.
 * First line starts with `!update` 
 * Second line will the index of event in which you are going to add your attendee. i.e `5`
 * Third line will contain the email address of attendees seperated by spaces. i.e `xyz@gmai.com abx@gmail.com pqr@work.ac.in`


 
### Running the bot
 * Make the following *. json* file that holds the bot key:

 `{ 'discord': 'FAKE000API000KEY000'}`

 * Run it:

 `python SchedulerBot`

### Examples
![SchedulerBotExamples](http://i.imgur.com/99wAUjN.png)



