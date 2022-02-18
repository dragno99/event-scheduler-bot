# Event scheduler bot for discord
A simple Discord bot created in golang that allows admin of a server to create scheduled google calender events. This bot will help admin to schedule an event with auto generated google's meet link and add attendees and feed this event to attendees's google calender.


### Functions of bot
 * Create and schedule an event.
 * Add attendees in an already scheduled event.
 * Show upcoming events.
 * Delete an event.

### Create and schedule an event
 * To create an event , input data will be in the following format.
 * First line starts with `!schedule` 
 * Second line will contain the summary of event.
 * Third line will contain the start date and time of event. `YYYY-MM-DDTHH-MM-SS` , i.e `2022-02-22T09-00-00`

### Running the bot
 * Make the following *. json* file that holds the bot key:

 `{ 'discord': 'FAKE000API000KEY000'}`

 * Run it:

 `python SchedulerBot`

### Examples
![SchedulerBotExamples](http://i.imgur.com/99wAUjN.png)



