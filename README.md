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
 * Third line will contain the start date and time of event. `YYYY-MM-DDTHH-MM-SS` , i.e `2022-02-22T09:00:00`
 * Fourth line will contain the end date and time of event. `YYYY-MM-DDTHH-MM-SS` , i.e `2022-02-22T11:00:00`
 * Fifth line will contain the email address of attendees seperated by spaces. i.e `xyz@gmai.com abx@gmail.com pqr@work.ac.in`
 * ![image](https://user-images.githubusercontent.com/76701875/154680826-66aff4e4-9eea-4e65-b005-32eb50a98936.png)


### Show upcoming events.
 * To see upcoming events, enter command `!upcoming`
 * ![image](https://user-images.githubusercontent.com/76701875/154680929-4bc6ee3c-aa61-4b84-985a-2af2103b3234.png)


### Add attendees in an already scheduled event.
 * To add attendees in an already scheduled event , first you need to see upcoming event by the `!upcoming` command.
 * After getting the list of upcoming event, You can add any number of attendes by the following command.
 * First line starts with `!update` 
 * Second line will contian the index of event in which you are going to add your new attendee. i.e `5`
 * Third line will contain the email address of attendees seperated by spaces. i.e `xyz@gmai.com abx@gmail.com pqr@work.ac.in`
 * ![image](https://user-images.githubusercontent.com/76701875/154681345-a53a6af2-6395-4435-abc8-35cc6b15910f.png)


### Delete an event.
 * To Delete an event,, first you need to see upcoming events by the `!upcoming` command.
 * After getting the list of upcoming events, now you can delete any event from the list.
 * First line starts with `!delete` 
 * Second line will contian the index of event which you want to delete. i.e `5`
 * ![image](https://user-images.githubusercontent.com/76701875/154681457-3b2dff6f-a49f-4bb8-b5ad-59322556de2d.png)





