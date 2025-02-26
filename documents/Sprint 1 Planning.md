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


### User Story #1.1

As a user, I want to be able to create an account with an email, real name, and password.


<table>
  <tr>
   <td>#</td>
   <td>Description</td>
   <td>Estimated Time</td>
   <td>Owner(s)</td>
  </tr>
  <tr>
   <td>1</td>
   <td>Create API endpoints for registration</td>
   <td>2 hr</td>
   <td>Cate</td>
  </tr>
  <tr>
   <td>2</td>
   <td>Create “Register New Account” UI</td>
   <td>2 hr</td>
   <td>Cate</td>
  </tr>
  <tr>
   <td>3</td>
   <td>Write tests for creating a new account, and creating duplicate accounts</td>
   <td>2 hr</td>
   <td>Cate</td>
  </tr>
  <tr>
   <td>4</td>
   <td>Implement checking for duplicate accounts during registration</td>
   <td>2 hr</td>
   <td>Cate</td>
  </tr>
  <tr>
   <td>5</td>
   <td>Implement password strength requirements.</td>
   <td>2 hr</td>
   <td>Cate</td>
  </tr>
</table>


Acceptance Criteria:



1. Given that an email is already registered, it cannot be registered again for a different account and an error message will pop up.
2. Given that the user does not have an account, they can create an account with an unregistered email.
3. Given that the user has successfully registered, they should be directed to their home screen.
4. Given that an email is already registered, users can navigate back to the log-in screen.
5. Given that a user enters a password without capital letters, lowercase letters, special characters, and below a specific length, it will not be accepted during registration.


### User Story #1.2

As a user, I want to be able to log in with an email and password so that I can see information for me and my team.


<table>
  <tr>
   <td>#</td>
   <td>Description</td>
   <td>Estimated Time</td>
   <td>Owner(s)</td>
  </tr>
  <tr>
   <td>1</td>
   <td>Implement login verification</td>
   <td>2 hr</td>
   <td>Cate</td>
  </tr>
  <tr>
   <td>2</td>
   <td>Implement error messages for invalid attempts</td>
   <td>2 hr</td>
   <td>Cate</td>
  </tr>
  <tr>
   <td>3</td>
   <td>Write tests for logging in</td>
   <td>2 hr</td>
   <td>Cate</td>
  </tr>
  <tr>
   <td>4</td>
   <td>Create “Log-In” UI</td>
   <td>2 hr</td>
   <td>Cate</td>
  </tr>
  <tr>
   <td>5</td>
   <td>Create API endpoints for login</td>
   <td>2 hr</td>
   <td>Cate</td>
  </tr>
  <tr>
   <td>6</td>
   <td>Implement password hashing</td>
   <td>3 hr</td>
   <td>Cate</td>
  </tr>
  <tr>
   <td>7</td>
   <td>Implement salting passwords</td>
   <td>3 hr</td>
   <td>Cate</td>
  </tr>
  <tr>
   <td>8</td>
   <td>Write tests to ensure password safety and login functionality</td>
   <td>2 hr</td>
   <td>Cate</td>
  </tr>
</table>


Acceptance Criteria:



1. Given that an email is registered, the account can be logged in to.
2. Given that the user is successfully logged in, they should be directed to their home screen.
3. Given that the user submits invalid login data, an error message will pop up.


### User Story #2.1

As a user, I want to be able to create a group so that I can begin my project


<table>
  <tr>
   <td>#</td>
   <td>Description</td>
   <td>Estimated Time</td>
   <td>Owner(s)</td>
  </tr>
  <tr>
   <td>1</td>
   <td>Set up mongoDB tables for groups</td>
   <td>2 Hr</td>
   <td>John</td>
  </tr>
  <tr>
   <td>2</td>
   <td>Create /projects/new API endpoint</td>
   <td>5 Hr</td>
   <td>John</td>
  </tr>
  <tr>
   <td>3</td>
   <td>Write tests for API endpoint</td>
   <td>3 Hr</td>
   <td>John</td>
  </tr>
  <tr>
   <td>4</td>
   <td>UI Modal for creating groups</td>
   <td>3 Hr</td>
   <td>John</td>
  </tr>
  <tr>
   <td>5</td>
   <td>Write tests for group creation modal</td>
   <td>1 Hr</td>
   <td>John</td>
  </tr>
</table>


Acceptance Criteria:



1. Given that the project creation modal is set up, when the user presses the create project button, a form should pop up with the required fields for project creation.
2. Given that the user isn’t logged in, when the user submits the create project form, the request should be rejected.
3. Given that the user submits invalid information in the project creation form, when the user submits the data, the server should return and the UI should display a helpful error message.
4. Given that the user submits valid information, when the user visits the user home page, the project should appear. 


### User Story #2.2

As a user I want to generate an invite link to the group so that I can collaborate with other students


<table>
  <tr>
   <td>#</td>
   <td>Description</td>
   <td>Estimated Time</td>
   <td>Owner(s)</td>
  </tr>
  <tr>
   <td>1</td>
   <td>Create an ‘/invite’ route</td>
   <td>4 Hr</td>
   <td>Ethan</td>
  </tr>
  <tr>
   <td>2</td>
   <td>Create tests for ‘/invite’ route</td>
   <td>1 Hr</td>
   <td>Ethan</td>
  </tr>
  <tr>
   <td>3</td>
   <td>Make a QR code for other users to scan on their phones to join the project</td>
   <td>2 Hr</td>
   <td>Ethan</td>
  </tr>
  <tr>
   <td>4</td>
   <td>Create UI to invite members to project</td>
   <td>0.5 Hr</td>
   <td>Ethan</td>
  </tr>
</table>


Acceptance Criteria:



1. Given the user is logged in, when the ‘invite’ route is visited, it will return a token expiring in a week that can be used to invite new members
2. Given the user is logged out, when the ‘/invite’ route is visited, it will return a 401 unauthorized error
3. Given the user is logged in, when they click the share button, the application should contact the server, get the token, and display a URL and QR code ot join


### User Story #2.4

As a user I want to click on invite links and automatically be added to that group


<table>
  <tr>
   <td>#</td>
   <td>Description</td>
   <td>Estimated Time</td>
   <td>Owner(s)</td>
  </tr>
  <tr>
   <td>1</td>
   <td>Implement ‘/join’ route which takes join session token as parameter and adds authenticated user to project</td>
   <td>3 Hr</td>
   <td>Ethan</td>
  </tr>
  <tr>
   <td>2</td>
   <td>Write tests for ‘/join’ route</td>
   <td>1 hr</td>
   <td>Ethan</td>
  </tr>
  <tr>
   <td>3</td>
   <td>Write UI for invite accept/decline</td>
   <td>2 Hr</td>
   <td>Ethan</td>
  </tr>
  <tr>
   <td>4</td>
   <td>Write UI for invalid invitation</td>
   <td>1 Hr</td>
   <td>Ethan</td>
  </tr>
  <tr>
   <td>5</td>
   <td>Write UI for expired invitation</td>
   <td>1 Hr</td>
   <td>Ethan</td>
  </tr>
  <tr>
   <td>6</td>
   <td>Write tests for UI</td>
   <td>1 Hr</td>
   <td>Ethan</td>
  </tr>
</table>


Acceptance Criteria:



1. Given an expired invite link, when the user visits it, they will see a notice that it’s expired
2. Using a link while signed out will redirect to login page (user story 2.5 will improve the flow)
3. Given a valid invite link, when the user visits it, they will see a confirmation page with basic project info and an option to accept or decline
4. Given an invalid invite link, when the user visits it, they will see a message indicating the link is invalid


### User Story #2.6

As a user I want to be able to leave a group.


<table>
  <tr>
   <td>#</td>
   <td>Description</td>
   <td>Estimated Time</td>
   <td>Owner(s)</td>
  </tr>
  <tr>
   <td>1</td>
   <td>Implement a `leave` endpoint in the websocket connection that removes a user from the group</td>
   <td>3 Hr</td>
   <td>Henry</td>
  </tr>
  <tr>
   <td>2</td>
   <td>Integrate a “Leave group” button into the UI as well as an “Are you sure?” modal</td>
   <td>3 hr</td>
   <td>Henry</td>
  </tr>
  <tr>
   <td>3</td>
   <td>Call the endpoint upon the user confirming leaving the group and redirect to the main page</td>
   <td>0.5 Hr</td>
   <td>Henry</td>
  </tr>
  <tr>
   <td>4</td>
   <td>Write integration tests for the user flow</td>
   <td>2 Hr</td>
   <td>Henry</td>
  </tr>
</table>


Acceptance Criteria:



1. Upon clicking the “leave” button, a confirmation modal will appear and clicking “Confirm” will cause the user to leave the group and redirect to the homepage. The group will not appear there.
2. Upon clicking the “leave” button, a confirmation modal will appear and clicking “Cancel” will cause the modal to close and nothing to happen.
3. An invitation link will allow the user to rejoin the group


### User Story #2.7

As a user I want to be able to delete a group.


<table>
  <tr>
   <td>#</td>
   <td>Description</td>
   <td>Estimated Time</td>
   <td>Owner(s)</td>
  </tr>
  <tr>
   <td>1</td>
   <td>Implement a `delete` endpoint in the websocket connection that deletes a group from the database along with all invitation links.</td>
   <td>3 Hr</td>
   <td>Henry</td>
  </tr>
  <tr>
   <td>2</td>
   <td>Integrate a “Delete group” button into the UI as well as an “Are you sure?” modal. It will be invisible or disabled when other people have not left the group.</td>
   <td>1 hr</td>
   <td>Henry</td>
  </tr>
  <tr>
   <td>3</td>
   <td>Call the endpoint upon the user confirming deleting the group and redirect to the main page</td>
   <td>0.5 Hr</td>
   <td>Henry</td>
  </tr>
  <tr>
   <td>4</td>
   <td>Write integration tests for the user flow</td>
   <td>2 Hr</td>
   <td>Henry</td>
  </tr>
</table>


Acceptance Criteria:



1. Upon clicking the “delete” button, a confirmation modal will appear and clicking “Confirm” will cause the group to be deleted and the user redirected to the main page. The group will not appear there.
2. Upon clicking the “delete” button, a confirmation modal will appear and clicking “Cancel” will cause the modal to close and nothing to happen.
3. Invitation links will not work after deleting the group.


### User Story #2.8

As a user, I want to be able to see a list of groups I am in so that I can choose which group I want to view


<table>
  <tr>
   <td>#</td>
   <td>Description</td>
   <td>Estimated Time</td>
   <td>Owner(s)</td>
  </tr>
  <tr>
   <td>1</td>
   <td>Create /projects API endpoint</td>
   <td>2 Hr</td>
   <td>John</td>
  </tr>
  <tr>
   <td>2</td>
   <td>Write tests for API endpoint</td>
   <td>2 Hr</td>
   <td>John</td>
  </tr>
  <tr>
   <td>3</td>
   <td>Create user home page UI</td>
   <td>5 Hr</td>
   <td>John</td>
  </tr>
</table>


Acceptance Criteria:



1. Given that the UI has been created, when the user navigates to the user home page, they should see a grid of projects that they are a part of.
2. Given that the UI has been created, when the user navigates to the user home page, they should see a button to create a new project.
3. Given that the user is part of at least 1 project, when the user clicks on a project’s card, it should navigate them to that project’s home page.
4. Given that the user is not logged in, when the user navigates to the user home page, they should be redirected to the login screen.


### User Story #2.10

Create and view a group information panel with a project name, project description, and links in the description automatically converted to anchor tags


<table>
  <tr>
   <td>#</td>
   <td>Description</td>
   <td>Estimated Time</td>
   <td>Owner(s)</td>
  </tr>
  <tr>
   <td>1</td>
   <td>Create project page UI</td>
   <td>3 Hr</td>
   <td>John</td>
  </tr>
  <tr>
   <td>2</td>
   <td>Write tests for displaying project page</td>
   <td>1 Hr</td>
   <td>John</td>
  </tr>
  <tr>
   <td>3</td>
   <td>Load state from database</td>
   <td>3 Hr</td>
   <td>John</td>
  </tr>
  <tr>
   <td>4</td>
   <td>Automatically turn links into anchor tags</td>
   <td>1 Hr</td>
   <td>John</td>
  </tr>
</table>


Acceptance Criteria:



1. Given that the project exists, when the user navigates to the project page, they see the project name.
2. Given that the project exists, when the user navigates to the project page, they see the project description.
3. Given that the project description contains links, when the user navigates to the project page, they see the links in the description as anchor tags.
4. Given that the user is not logged in, when the user navigates to the project page, they should be redirected to the login screen.


### User Story #5.2

As a user I want to be able to manually set availabilities if I don’t use a traditional calendar


<table>
  <tr>
   <td>#</td>
   <td>Description</td>
   <td>Estimated Time</td>
   <td>Owner(s)</td>
  </tr>
  <tr>
   <td>1</td>
   <td>Create availability input component</td>
   <td>4 Hr</td>
   <td>Ethan</td>
  </tr>
  <tr>
   <td>2</td>
   <td>Save availability data to backend</td>
   <td>2 Hr</td>
   <td>Ethan</td>
  </tr>
  <tr>
   <td>3</td>
   <td>Translate availabilities across time zones</td>
   <td>1 Hr</td>
   <td>Ethan</td>
  </tr>
  <tr>
   <td>4</td>
   <td>Write unit tests for inputting availability</td>
   <td>1 Hr</td>
   <td>Ethan</td>
  </tr>
  <tr>
   <td>5</td>
   <td>Write unit test for saving availability to server</td>
   <td>1 Hr</td>
   <td>Ethan</td>
  </tr>
  <tr>
   <td>6</td>
   <td>Write integration tests component saving to server</td>
   <td>1 Hr</td>
   <td>Ethan</td>
  </tr>
</table>


Acceptance Criteria:



1. Given the UI is implemented, when the user navigates to the availability page, they should see a table. On the X axis are dates, the Y axis is times in 30-minute increments, and the cells can be checked and unchecked to indicate whether or not they are available during that time block.
2. Given the UI is implemented, when a user navigates to the availability page, they should see a drop-down to change their time zone
3. Given time zones work correctly, when the user switches their timezone, the times and dates should shift to reflect their current time zone


### User Story #5.3

As a user I want to be able to view everyone’s availability as a calendar in a weekly view


<table>
  <tr>
   <td>#</td>
   <td>Description</td>
   <td>Estimated Time</td>
   <td>Owner(s)</td>
  </tr>
  <tr>
   <td>1</td>
   <td>Extend availability component to combine availabilities</td>
   <td>2 Hr</td>
   <td>Ethan</td>
  </tr>
  <tr>
   <td>2</td>
   <td>Add a tooltip to availability component</td>
   <td>2 Hr</td>
   <td>Ethan</td>
  </tr>
  <tr>
   <td>3</td>
   <td>Write tests to ensure rendering correctly</td>
   <td>1 Hr</td>
   <td>Ethan</td>
  </tr>
</table>


Note: Non-User Story Task #1 is a subtask of this because networking is required for this to function correctly.

Acceptance Criteria:



1. Given the group calendar works, when multiple different people input their availability, the combined availability calendar should be darker in areas where the most people are available
2. Given the group calendar works, when hovering over a time block, the user should see a list of members who are available or not during that time block
3. Given the group calendar works, as you or other members input their availability in real-time, changes should be reflected immediately


## Non-User Story Tasks


### Task #1

Client-Server communication will be done using WebSockets to allow real time updates


<table>
  <tr>
   <td>#</td>
   <td>Description</td>
   <td>Estimated Time</td>
   <td>Owner(s)</td>
  </tr>
  <tr>
   <td>1</td>
   <td>Create a class to encapsulate the state and the WebSocket connection along with a meaningful Typescript API for the class including local-only stub implementations. Also set up the project page for Svelte so that it uses this class as a parameter.</td>
   <td>3 Hr</td>
   <td>Henry</td>
  </tr>
  <tr>
   <td>2</td>
   <td>Establish a WebSocket connection endpoint on the server and allow the webpage to connect to it</td>
   <td>5 Hr</td>
   <td>Henry</td>
  </tr>
  <tr>
   <td>3</td>
   <td>Set up the Svelte interface to accept the program state and make the state update on changes broadcast by the server</td>
   <td>3 Hr</td>
   <td>Henry</td>
  </tr>
  <tr>
   <td>4</td>
   <td>Establish an API on the server to handle updating state on events being received</td>
   <td>3 Hr</td>
   <td>Henry</td>
  </tr>
  <tr>
   <td>5</td>
   <td>Create a temporary demonstration of the project page including a radio button that disables on click and an un-disable button to demonstrate networking</td>
   <td>0.5 Hr</td>
   <td>Henry</td>
  </tr>
  <tr>
   <td>6</td>
   <td>Implement `appendInProject`, `updateInProject`, and `deleteInProject` events on the server and client</td>
   <td>5 Hr</td>
   <td>Henry</td>
  </tr>
  <tr>
   <td>7</td>
   <td>As necessary, refactor my teammate’s code to use the new events</td>
   <td>2 Hr</td>
   <td>Henry</td>
  </tr>
</table>


Acceptance Criteria:



1. Selecting a radio option on one machine will automatically cause it to select on the other
2. Clicking the un-disable button will deselect the option from the radio button and re-enable it on both machines
3. Clicking different options on two machines will cause the same option to be selected on both machines. Which option is unspecified, the important part is that it’s the same one
4. Any project specific UI implemented by teammates should reflect on other machines


### 


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
    6. ~~Create a group so that I can begin my project~~
    7. ~~Generate an invite link to the group so that I can collaborate with other students~~
    8. (If time allows) Invite people directly in the app using their email
    9. ~~Click on invite links and automatically be added to that group~~
    10. Be directed to the login page when clicking an invite link if not signed in and automatically being added to the group after signing in
    11. ~~Leave a group if I am leaving that group or changing group~~
    12. ~~Delete a group so that data can be deleted if wanted at the end of the project~~
    13. ~~See a list of groups I am in so that I can choose which group I want to view~~
    14. View my teammates’ profiles, specifically names, profile pictures, and contact information
    15. ~~Create and view a group information panel with a project name, project description, and links in the description automatically converted to anchor tags~~
3. **Tasks**
    16. View current tasks in a simplified 3 column kanban board
    17. View who is assigned to each task
    18. Create tasks so that I can track those tasks throughout the lifetime of the group project
    19. Assign due dates to tasks
    20. Assign time estimates to tasks
    21. Delete tasks and mark tasks as completed
    22. Assign/unassign tasks so that the entire team knows who is responsible for which tasks
    23. Track time spent on tasks so that the team knows how much time has been spent by who on each task
    24. (if time allows) View a personal window with all tasks I am assigned to across multiple groups
4. **Stats**
    25. See a burndown chart so that I can track the team’s progress toward our goals
    26. Split the burndown chart into a stacked line plot by user
    27. Show the burndown chart by estimated time
    28. See a pie chart with how much time each team member has recorded
5. **Calendar**
    29. (if time allows) Upload my calendar from another program into Pleiades so that it can be used for finding team availabilities
    30. ~~Manually set availabilities if I don’t use a traditional calendar~~
    31. ~~View everyone’s availability as a calendar in a weekly view~~
    32. View due dates of tasks in the calendar view in a monthly view
    33. See a color coded view of the calendar/tasks indicative of the progress made
6. **Notifications**
    34. (if time allows) View notifications in-app of updates and changes made to the project
    35. (if time allows) Be able to change/customize my notification settings 
7. **Voting**
    36. Start a vote to decide project direction when consensus cannot be reached
    37. Vote in a project vote
    38. Display the results of of a vote


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
