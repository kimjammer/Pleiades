# **Team 5 Sprint 1 Planning**

**Pleiades**

Cate Harrison, John Kim, Henry Rovnyak, Ethan Dawes


## Sprint Overview

In this first sprint, we will be focusing on building up the infrastructure and codebase necessary to build features on top of. We will need to coordinate how common patterns in both the front and backend should be implemented, and make sure the project structure is set up in a meaningful way. We will also start implementing the core features of the application, as well as conducting a user study to determine our final UI layout.

**Scrum Master:** Ethan Dawes

**Meeting Plan:** Wednesday 11:30 AM, Thursday 8:50 AM


### Risks and Challenges

Since there is no codebase to start off from, this initial period of programming is the most prone to bad code patterns, incorrect implementations, and merge conflicts. We will take special care to ensure clear and consistent communication to help alleviate these symptoms. Another challenge is that the core code we write in this sprint will be depended upon by future features, so ensuring bug-free code in these stories will lower debugging and refactoring workloads in the future.




## Current Sprint Detail


### User Story #2.1

As a user, I want to be able to create a group so that I can begin my project


| # | Description                          | Estimated Time | Owner(s) |
|---|--------------------------------------|----------------|----------|
| 1 | Set up mongoDB tables for groups     | 2 Hr           | John     |
| 2 | Create /projects/new API endpoint    | 5 Hr           | John     |
| 3 | Write tests for API endpoint         | 3 Hr           | John     |
| 4 | UI Modal for creating groups         | 4 Hr           | John     |
| 5 | Write tests for group creation modal | 1 Hr           | John     |



Acceptance Criteria:



1. Given that the project creation modal is set up, when the user presses the create project button, a form should pop up with the required fields for project creation.
2. Given that the API endpoint has been created, when the user submits the create project form, the data should be validated and saved to the database.
3. Given that the user isn’t logged in, when the user submits the create project form, the request should be rejected.
4. Given that the user submits invalid information in the project creation form, when the user submits the data, the server should return and the UI should display a helpful error message.


### User Story #2.8

As a user, I want to be able to see a list of groups I am in so that I can choose which group I want to view


| # | Description                   | Estimated Time | Owner(s) |
|---|-------------------------------|----------------|----------|
| 1 | Create /projects API endpoint | 2 Hr           | John     |
| 2 | Write tests for API endpoint  | 2 Hr           | John     |
| 3 | Create user home page UI      | 5 Hr           | John     |



Acceptance Criteria:



1. Given that the UI has been created, when the user navigates to the user home page, they should see a grid of projects that they are a part of.
2. Given that the UI has been created, when the user navigates to the user home page, they should see a button to create a new project.
3. Given that the user is part of at least 1 project, when the user clicks on a project’s card, it should navigate them to that project’s home page.
4. Given that the user is not logged in, when the user navigates to the user home page, they should be redirected to the login screen.


### User Story #1.1

As a user, I want to be able to create an account with an email, real name, and password.


| # | Description                                                             | Estimated Time | Owner(s) |
|---|-------------------------------------------------------------------------|----------------|----------|
| 1 | Create API endpoints for registration                                   | 2 hr           | Cate     |
| 2 | Create “Register New Account” UI                                        | 2 hr           | Cate     |
| 3 | Write tests for creating a new account, and creating duplicate accounts | 2 hr           | Cate     |
| 4 | Implement checking for duplicate accounts during registration           | 2 hr           | Cate     |
| 5 | Implement password strength requirements.                               | 2 hr           | Cate     |



Note: Non-User Story Task #1 is a subtask of this because both the frontend and backend must be deployed and functioning to be able to access and log into the application.

Acceptance Criteria:



1. Given that an email is already registered, it cannot be registered again for a different account and an error message will pop up.
2. Given that the user does not have an account, they can create an account with an unregistered email.
3. Given that the user has successfully registered, they should be directed to their home screen.
4. Given that an email is already registered, users can navigate back to the log-in screen.
5. Given that a user enters a password without capital letters, lowercase letters, special characters, and below a specific length, it will not be accepted during registration.


### User Story #1.2

As a user, I want to be able to log in with an email and password so that I can see information for me and my team.


| # | Description                                                   | Estimated Time | Owner(s) |
|---|---------------------------------------------------------------|----------------|----------|
| 1 | Implement login verification                                  | 2 hr           | Cate     |
| 2 | Implement error messages for invalid attempts                 | 2 hr           | Cate     |
| 3 | Write tests for logging in                                    | 2 hr           | Cate     |
| 4 | Create “Log-In” UI                                            | 2 hr           | Cate     |
| 5 | Create API endpoints for login                                | 2 hr           | Cate     |
| 6 | Implement password hashing                                    | 3 hr           | Cate     |
| 7 | Implement salting passwords                                   | 3 hr           | Cate     |
| 8 | Write tests to ensure password safety and login functionality | 3 hr           | Cate     |



Note: Non-User Story Task #3 is a subtask of this because generating authentication tokens is the final step of logging in from the perspective of the server.

Acceptance Criteria:



1. Given that an email is registered, the account can be logged in to.
2. Given that the user is successfully logged in, they should be directed to their home screen.
3. Given that the user submits invalid login data, an error message will pop up.
4. Given that an account is registered, the true password will not be stored on the server.
5. Given that a user has a password, a unique string will be generated to salt their password before hashing.
6. Given that a user has an account, a reasonably strong hashing method will be used.


### User Story #5.2

As a user I want to be able to manually set availabilities if I don’t use a traditional calendar


| # | Description                                        | Estimated Time | Owner(s) |
|---|----------------------------------------------------|----------------|----------|
| 1 | Create availability input component                | 4 Hr           | Ethan    |
| 2 | Save availability data to backend                  | 2 Hr           | Ethan    |
| 3 | Translate availabilities across time zones         | 1 Hr           | Ethan    |
| 4 | Write unit tests for inputting availability        | 1 Hr           | Ethan    |
| 5 | Write unit test for saving availability to server  | 1 Hr           | Ethan    |
| 6 | Write integration tests component saving to server | 1 Hr           | Ethan    |



Acceptance Criteria:



1. Given the UI is implemented, when the user navigates to the availability page, they should see a table. On the X axis are dates, the Y axis is times in 30-minute increments, and the cells can be checked and unchecked to indicate whether or not they are available during that time block.
2. Given the UI is implemented, when a user navigates to the availability page, they should see a drop-down to change their time zone
3. Given time zones work correctly, when the user switches their timezone, the times and dates should shift to reflect their current time zone


### User Story #5.3

As a user I want to be able to view everyone’s availability as a calendar in a weekly view


| # | Description                                             | Estimated Time | Owner(s) |
|---|---------------------------------------------------------|----------------|----------|
| 1 | Extend availability component to combine availabilities | 2 Hr           | Ethan    |
| 2 | Add a tooltip to availability component                 | 2 Hr           | Ethan    |
| 3 | Write tests to ensure rendering correctly               | 1 Hr           | Ethan    |



Note: Non-User Story Task #2 is a subtask of this because networking is required for this to function correctly.

Acceptance Criteria:



1. Given the group calendar works, when multiple different people input their availability, the combined availability calendar should be darker in areas where the most people are available
2. Given the group calendar works, when hovering over a time block, the user should see a list of members who are available or not during that time block
3. Given the group calendar works, as you or other members input their availability in real-time, changes should be reflected immediately


### User Story #2.2

As a user I want to generate an invite link to the group so that I can collaborate with other students


| # | Description                                                                | Estimated Time | Owner(s) |
|---|----------------------------------------------------------------------------|----------------|----------|
| 1 | Create an ‘/invite’ route                                                  | 4 Hr           | Ethan    |
| 2 | Create tests for ‘/invite’ route                                           | 1 Hr           | Ethan    |
| 3 | Make a QR code for other users to scan on their phones to join the project | 2 Hr           | Ethan    |
| 4 | Create UI to invite members to project                                     | 0.5 Hr         | Ethan    |



Acceptance Criteria:



1. Given the user is logged in, when the ‘invite’ route is visited, it will return a token expiring in a week that can be used to invite new members
2. Given the user is logged out, when the ‘/invite’ route is visited, it will return a 401 unauthorized error
3. Given the user is logged in, when they click the share button, the application should contact the server, get the token, and display a URL and QR code ot join


### User Story #2.4

As a user I want to click on invite links and automatically be added to that group


| # | Description                                                                                                | Estimated Time | Owner(s) |
|---|------------------------------------------------------------------------------------------------------------|----------------|----------|
| 1 | Implement ‘/join’ route which takes join session token as parameter and adds authenticated user to project | 3 Hr           | Ethan    |
| 2 | Write tests for ‘/join’ route                                                                              | 1 hr           | Ethan    |
| 3 | Write UI for invite accept/decline                                                                         | 2 Hr           | Ethan    |
| 4 | Write UI for invalid invitation                                                                            | 1 Hr           | Ethan    |
| 5 | Write UI for expired invitation                                                                            | 1 Hr           | Ethan    |
| 6 | Write tests for UI                                                                                         | 1 Hr           | Ethan    |



Acceptance Criteria:



1. Given an expired invite link, when the user visits it, they will see a notice that it’s expired
2. Using a link while signed out will redirect to login page (task 2.5 will improve the flow)
3. Given a valid invite link, when the user visits it, they will see a confirmation page with basic project info and an option to accept or decline
4. Given an invalid invite link, when the user visits it, they will see a message indicating the link is invalid


## Non-User Story Tasks


### Task #1

As a developer, we want to have CI/CD so that deployments can be made quickly to the prod server.


| # | Description                                    | Estimated Time | Owner(s) |
|---|------------------------------------------------|----------------|----------|
| 1 | Setup automated website builds and deployments | 2 Hr           | John     |
| 2 | Containerize the backend application           | 2 Hr           | John     |
| 3 | Setup automated backend builds and deployments | 3 Hr           | John     |
| 4 | Test the CI/CD system                          | 0.5 Hr         | John     |



Acceptance Criteria:



1. Given the website CI/CD is set up, when a website commit is pushed, it should be automatically deployed within 30 minutes.
2. Given the backend CI/CD is set up, when a backend commit is pushed, it should be automatically deployed within 30 minutes.
3. Given the backend has been containerized, when code changes are made, they should be able to be built and run with only the docker compose file.


### Task #2

Client-Server communication will be done using WebSockets to allow real time updates


| # | Description                                                                                                                                                                                                                                             | Estimated Time | Owner(s) |
|---|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------|----------|
| 1 | Create a class to encapsulate the state and the WebSocket connection along with a meaningful Typescript API for the class including local-only stub implementations. Also set up the project page for Svelte so that it uses this class as a parameter. | 3 Hr           | Henry    |
| 2 | Establish a WebSocket connection endpoint on the server and allow the webpage to connect to it                                                                                                                                                          | 5 Hr           | Henry    |
| 3 | Set up the Svelte interface to accept the program state and make the state update on changes broadcast by the server                                                                                                                                    | 3 Hr           | Henry    |
| 4 | Establish an API on the server to handle updating state on events being received                                                                                                                                                                        | 3 Hr           | Henry    |
| 5 | Create a temporary demonstration of the project page including a radio button that disables on click and an un-disable button to demonstrate networking                                                                                                 | 0.5 Hr         | Henry    |
| 6 | Implement `appendInProject`, `updateInProject`, and `deleteInProject` events on the server and client                                                                                                                                                   | 5 Hr           | Henry    |
| 7 | As necessary, refactor my teammate’s code to use the new events                                                                                                                                                                                         | 2 Hr           | Henry    |



Acceptance Criteria:



1. Selecting a radio option on one machine will automatically cause it to select on the other
2. Clicking the un-disable button will deselect the option from the radio button and re-enable it on both machines
3. Clicking different options on two machines will cause the same option to be selected on both machines. Which option is unspecified, the important part is that it’s the same one
4. Any project specific UI implemented by teammates should reflect on other machines


### Task #3

We need a system to authenticate endpoints using tokens


| # | Description                                                                                                                                     | Estimated Time | Owner(s) |
|---|-------------------------------------------------------------------------------------------------------------------------------------------------|----------------|----------|
| 1 | Create stub functions for token creation and verification that do not use cryptography (to allow others to integrate them into their endpoints) | 0.5 Hr         | Henry    |
| 2 | Implement the function that creates tokens containing the user’s ID, creation date, and expiration date using a message authentication code     | 5 Hr           | Henry    |
| 3 | Implement the function that decodes tokens and verifies the MAC                                                                                 | 3 Hr           | Henry    |
| 4 | Create tests to verify correct encoding, decoding, and validation of tests                                                                      | 2 Hr           | Henry    |
| 4 | As necessary, assist other teammates in integrating the tokens                                                                                  | 1 Hr           | Henry    |



Acceptance Criteria:



1. For all authenticated endpoints, a forged token will not grant access to the endpoint.
2. For all authenticated endpoints, a true token *will* grant access to the endpoint
3. For all authenticated endpoints, the endpoint will be able to identify who is accessing the endpoint




### Remaining Backlog


### Functional

As a user, I want to be able to



1. **Account**
    1. ~~Create an account with a email, real name, and password~~
    2. ~~Log in with a email and password so that I can see information for me and my team~~
    3. (If time allows) login with social accounts (Google)
    4. (If time allows) Reset my password so that I can recover my account
    5. Upload a profile picture
2. **Group**
    1. ~~Create a group so that I can begin my project~~
    2. ~~Generate an invite link to the group so that I can collaborate with other students~~
    3. (If time allows) Invite people directly in the app using their email
    4. ~~Click on invite links and automatically be added to that group~~
    5. Be directed to the login page when clicking an invite link if not signed in and automatically being added to the group after signing in
    6. Leave a group if I am leaving that group or changing group
    7. Delete a group so that data can be deleted if wanted at the end of the project
    8. ~~See a list of groups I am in so that I can choose which group I want to view~~
    9. View my teammates’ profiles, specifically names, profile pictures, and contact information
    10. Create and view a group information panel with a project name, project description, and links in the description automatically converted to anchor tags
3. **Tasks**
    1. View current tasks in a simplified 3 column kanban board
    2. View who is assigned to each task
    3. Create tasks so that I can track those tasks throughout the lifetime of the group project
    4. Assign due dates to tasks
    5. Assign time estimates to tasks
    6. Delete tasks and mark tasks as completed
    7. Assign/unassign tasks so that the entire team knows who is responsible for which tasks
    8. Track time spent on tasks so that the team knows how much time has been spent by who on each task
    9. (if time allows) View a personal window with all tasks I am assigned to across multiple groups
4. **Stats**
    1. See a burndown chart so that I can track the team’s progress toward our goals
    2. Split the burndown chart into a stacked line plot by user
    3. Show the burndown chart by estimated time
    4. See a pie chart with how much time each team member has recorded
5. **Calendar**
    1. (if time allows) Upload my calendar from another program into Pleiades so that it can be used for finding team availabilities
    2. ~~Manually set availabilities if I don’t use a traditional calendar~~
    3. ~~View everyone’s availability as a calendar in a weekly view~~
    4. View due dates of tasks in the calendar view in a monthly view
    5. See a color coded view of the calendar/tasks indicative of the progress made
6. **Notifications**
    1. (if time allows) View notifications in-app of updates and changes made to the project
    2. (if time allows) Be able to change/customize my notification settings 
7. **Voting**
    1. Start a vote to decide project direction when consensus cannot be reached
    2. Vote in a project vote
    3. Display the results of of a vote


### Non-functional



* Architecture
    * ~~Frontend written using the Svelte framework~~
    * ~~Utilize components from shadcn-svelte~~
    * ~~Backend written using Go~~
    * ~~Use client side rendering to make the backend and frontend independent~~
    * ~~Client-server communication will be done using websockets to allow real time updates~~
* Appearance
    * Use a minimalistic and modern style that is accessible and follows common, well known design patterns
* Security
    * ~~Use a rigorously tested library to perform password hashing and comply with security standards~~
    * Be cognizant to avoid and mitigate buffer overflows and XSS attacks 
* Performance and Usability
    * Animations must be smooth and we must ensure that we hit 60 fps on mobile level hardware
    * Initial page loading times must be under 1 second including client side rendering
    * We will perform user studies to ensure that our UI design is intuitive and understandable
    * (if time allows) We will make our design responsive to mobile devices
    * (if time allows) Allow using the website offline and sync when back online
* Developer experience
    * 60% code coverage
* ~~Have CI/CD so that deployments can be made quickly to the prod server~~
* John will host the deployment on his personal server
* We shall reify an impressive, useful product
* We shall get a good grade
