# Team 5 Sprint 2 Planning
**Pleiades**  
Cate Harrison, John Kim, Henry Rovnyak, Ethan Dawes

## Sprint Overview

In our second sprint, we will be implementing our application’s core features around tasks and time recording. As this is our app’s main selling point, we will need to put in effort to ensure that user flows are smooth and intuitive. Additionally, this sprint will put to use the realtime networking logic implemented in the last sprint so that realtime collaboration between team members is possible.

**Scrum Master:** John Kim

**Meeting Plan:** Wednesday 11:30 AM, Thursday 8:30 AM

### Risks and Challenges

This sprint will involve a high degree of collaboration, as the features we are implementing in this sprint are highly connected to each other. We hope that the use of frequent automated and manual tests will help us catch integration bugs early, before they compound on each other and create phenomena that are harder to debug. Another potential challenge is that most of these features need to synchronize in real time across different users, so we need to make sure that our UI state and backend database state stay consistent for all users, even as they connect, disconnect, or modify the project state.

## Current Sprint Detail

### User Story \#3.4

As a user, I want to assign due dates to tasks

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Create button and calendar component to select or modify due date | 2 hr | John |
| 2 | Create chip component to display due date on card | 1 hr | John |
| 3 | Convert date to UTC time based on local time | 0.5 hr | John |
| 4 | Write tests for due dates | 2 hr | John |
| 5 | Write backend test harness for e2e testing | 1 hr | John |

Acceptance Criteria:

1. Given that a task exists and has no due date, there is a button to input a due date.  
2. Given that a task exists and has a due date, there is a chip displaying the due date.  
3. Given that a task exists and has a due date, it can be changed by clicking on it.  
4. Given that a task has a due date, it is displayed in your local timezone.

### User Story \#3.5

As a user, I want to be able to assign time estimates to tasks

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Create button and duration component to select or modify the time estimate | 2 hr | John |
| 2 | Create chip component to display time estimate | 1 hr | John |
| 3 | Write tests for time estimates | 2 hr | John |

Acceptance Criteria:

1. Given that a task exists and has no time estimate, there is a button to add a time estimate.  
2. Given that a task exists and has a time estimate, there is a chip displaying the time estimate.  
3. Given that a task exists and has a time estimate, it can be changed by clicking on it.

### User Story \#3.8

As a user, I want to be able to track time spent on tasks so that the team knows how much time has been spent by who on each task

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Implement UI for timer | 1 hr | John |
| 2 | Implement server side logic for validating timer entries | 3 hr | John |
| 3 | UI for displaying how much time has been spent by who | 2 hr | John |
| 4 | Update time estimate chip to act as progress bar by spent time | 2 hr | John |
| 5 | Write tests for timer | 2 hr | John |

Acceptance Criteria:

1. Given a task, you can start a timer to record your work session.  
2. Given a timer is running, you can stop the timer to complete your work session.  
3. Given at least 1 user has at least 1 session for a task, you can see the total amount of time spent by each person on that task.  
4. Given that at least 1 session exists for a task, you can see the progress to the time estimate.

### User Story \#4.1

As a user, I want to be able to see a burndown chart so that I can track the team’s progress toward our goals.

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Implement burndown graph component | 3 hr | John |
| 2 | Write backend logic to aggregate time session data | 3 hr | John |
| 3 | Write tests for backend data aggregation | 2 hr | John |
| 4 | Write frontend tests for graph component | 1 hr | John |

Acceptance Criteria:

1. Given at least 1 task with both a time estimate and a session, you can see a burndown chart of total time estimate vs time spent on a daily interval.  
2. Given there is no task with both a time estimate and a session, the graph component shows a helpful “N/A” message to the user.  
3. Given a new time estimate or session is recorded, the graph updates in real time.

### User Story \#2.5

As a user, I want to be able to be directed to the login page when clicking an invite link if not signed in and automatically being added to the group after signing in

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Create automated testing for this flow | 1 hr | Ethan |
| 2 | Revert behavior in previous sprint, don't automatically redirect when signed out | 1 hr | Ethan |
| 3 | Redirect user to homepage with `?continue` url token | 0.5 hr | Ethan |
| 4 | On user homepage, implement project joining logic if `?continue` token detected in url  | 1.5 hr | Ethan |

Acceptance Criteria:

1. Given the user visits the /join page while signed out, they are not immediately directed to the landing page (as implemented last sprint). Instead, they see the join confirmation page.  
2. Given the user clicks the `accept` button, they are redirected to the login page  
3. Given the user did steps 1 & 2 of the acceptance criteria, after the user signs in, they are taken directly to the project page for the project they just accepted & become a project member.

### User Story \#2.9

As a user, I want to be able to view my teammates’ profiles, specifically names, profile pictures, and contact information

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Implement browser tests for showing teammate info | 1 hr | Ethan |
| 2 | Render teammates’ info | 2 hr | Ethan |

Acceptance Criteria:

1. Given I am looking at my teammate’s info tab, if I click the “contact” button, my email client opens to send them mail  
2. Given I am in the project, I will see a tab where the group info is displayed.  
3. Given I am looking at the teammate info tab, I will see each user’s names, profile pictures, and a method to contact them is displayed

### User Story \#3.3

As a user, I want to be able to create tasks so that I can track those tasks throughout the lifetime of the group project

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | button to create a new task | 0.5 hr | Ethan |
| 2 | modal with fields such as `title`, `description`, `due date`, `time estimate`, and `assignees`. | 2 hr | Ethan |
| 3 | Modal validation | 1 hr | Ethan |
| 4 | Natural language parsing from title | 3 hr | Ethan |
| 5 | Automated tests for each acceptance criteria | 2 hr | Ethan |
| 6 | Conduct user studies on what consumers want and what interface looks best | 2 hr | Ethan |

Acceptance Criteria:

1. Given the user is on the `tasks` tab of the project page, they will see a button to create a new task  
2. Given the user presses this button to create a new task, they are presented with a modal with fields such as `title`, `description`, `due date`, `time estimate`, and `assignees`.  
3. Given the user is looking at the task creation modal, they will not be able to create a task and will be presented with an error if they leave the title blank  
4. Given the user has typed a duration (5 hours/5hr) or a due date (next tuesday/Jan 15\) into the title field, the respective fields later in the form will be populated with this information

### User Story \#4.4

As a user, I want to be able to see a pie chart with how much time each team member has recorded

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Make the pie chart | 2 hr | Ethan |
| 2 | Write tests that the input data is correct | 1 hr | Ethan |

Acceptance Criteria:

1. Given I am looking at the designated tab on the project page and members have recorded time, there will be a pie chart,  where each pie slice represents the cumulative time a team member spent on any task in this project  
2. Given I  hover over a slice on the pie chart, a tooltip appears with  the user's name and how much time they spent  
3. Given I am looking at the designated tab on the project page and no members have recorded time, there will be a message indicating so and encouraging members to track time

### User Story \#5.4

As a user, I want to be able to view due dates of tasks in the calendar view in a monthly view

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Create a month-view calendar | 1 hr | Ethan |
| 2 | Place the tasks onto the calendar | 3 hr | Ethan |
| 3 | Implement month switching buttons | 2 hr | Ethan |
| 4 | Create tests for the calendar page | 2 hr | Ethan |
| 5 | Conduct user studies on what consumers want and what interface looks best | 2 hr | Ethan |

Acceptance Criteria:

1. Given I am looking at the designated tab on the project page, I will see a month-view calendar  
2. Given users have created tasks with due dates, each of those incomplete tasks titles is placed on its respective date.  
3. Given I press the left and right buttons at the top of the page, the month I am looking at changes

### 

### User Story \#3.1

As a user, I want to view current tasks in a simplified 3 column kanban board

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Implement the project mutation WS endpoints on the server  | 4 hr | Henry |
| 2 | Write tests for the websockets endpoints | 4 hr | Henry |
| 3 | Implement the 3 column kanban board | 3 hr | Henry |
| 4 | Display the tasks in the kanban board | 3 hr | Henry |
| 5 | Get Playwright working on Henry’s computer | 2 hr | Henry |
| 6 | Write tests for the kanban board | 2 hr | Henry |

Acceptance Criteria:

1. The project page will show the three column kanban board  
2. Each task will display as a card in the correct column  
3. Multiple connected users will display the same information

### User Story \#3.2

As a user, I want to view who is assigned to each task

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | For each card on the Kanban board, create a line displaying the owner if there is one and nothing if there is not | 1 hr | Henry |
| 2 | Write tests to ensure that the line is there when there is an assigned user and not when it is not there | 1 hr | Henry |

Acceptance Criteria:

1. If a task is not assigned to a user, the card should not display a user  
2. If a task is assigned a user, the card should display it in a line  
3. The user should be displayed as their full real name

### User Story \#3.6

As a user, I want to delete tasks and mark tasks as completed

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Implement drag and drop for the cards | 4 hr | Henry |
| 2 | Upon being dropped in a column, the card should be moved to that column | 1 hr | Henry |
| 3 | Override the right-click operation for cards with a custom tooltip | 3 hr | Henry |
| 4 | Add a “delete” button to the tooltip that will delete the card | 3 hr | Henry |

Acceptance Criteria:

1. Dragging and dropping cards will move the card to the column it was dropped in  
2. For a different user, the card will be moved to that column on their screen  
3. Right clicking a card will show a custom tooltip  
4. Clicking delete on the tooltip will delete the task  
5. The task will appear deleted on the screen of another user

### User Story \#1.5

As a user, I want to be able to upload a profile picture.

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Implement default profile pictures | 2 hr | Cate |
| 2 | Create UI to upload profile picture | 1 hr | Cate |
| 3 | Implement storage and display of profile picture | 3 hr | Cate |
| 4 | Implement security against non-image files | 3 hr | Cate |
| 5 | Write tests for image uploading | 1 hr | Cate |

Acceptance Criteria:

1. ### Given that the user uploads a valid image, it will be displayed as the profile picture

2. ### Given that the user tries to upload an invalid file, it will not work.

3. ### Given that a user has not uploaded a profile picture, they will have a default image instead.

### User Story \#3.7

As a user, I want to be able to assign/unassign tasks so that the entire team knows who is responsible for which tasks.

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Create UI elements to add users to tasks | 1 hr | Cate |
| 2 | Create UI elements to remove users from tasks | 1 hr | Cate |
| 3 | Create endpoints for adding and removing users from tasks | 4 hr | Cate |
| 4 | Create UI elements to display what users are assigned to which task | 1 hr | Cate |
| 5 | Create endpoint for fetching current assigned users for a task | 2 hr | Cate |
| 6 | Write tests for adding/removing users from tasks | 1 hr | Cate |

Acceptance Criteria:

1. ### Given that a user is not assigned to a task, they may join a task.

2. ### Given that a user is assigned to a task, they may leave a task.

3. ### Given that a user is already assigned to a task, they may not join it twice.

4. ### Given that there are users assigned to a task, all users in the workspace can see the current assigned users.

### User Story \#7.1

As a user, I want to be able to start a vote to decide project direction when consensus cannot be reached.

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Create UI for creating a vote | 2 hr | Cate |
| 2 | Create UI to display current polls | 2 hr | Cate |
| 3 | Create endpoint for creating a poll | 3 hr | Cate |
| 4 | Write tests for poll creation functionality | 3 hr | Cate |

Acceptance Criteria:

1. Given that a user is in a project, they can create a poll.  
2. Given that a poll exists in a project, users in that project can see that it exists.  
3. Given that a poll does not have at least a title and two options, it cannot be created.  
4. Given that a poll is created, it has a time at which it will end.

### Remaining Backlog

### Functional

As a user, I want to be able to

1. **Account**  
   1. ~~Create an account with a email, real name, and password~~  
   2. ~~Log in with a email and password so that I can see information for me and my team~~  
   3. (If time allows) login with social accounts (Google)  
   4. (If time allows) Reset my password so that I can recover my account  
   5. ~~Upload a profile picture~~  
2. **Group**  
   1. ~~Create a group so that I can begin my project~~  
   2. ~~Generate an invite link to the group so that I can collaborate with other students~~  
   3. (If time allows) Invite people directly in the app using their email  
   4. ~~Click on invite links and automatically be added to that group~~  
   5. ~~Be directed to the login page when clicking an invite link if not signed in and automatically being added to the group after signing in~~  
   6. ~~Leave a group if I am leaving that group or changing group~~  
   7. ~~Delete a group so that data can be deleted if wanted at the end of the project~~  
   8. ~~See a list of groups I am in so that I can choose which group I want to view~~  
   9. ~~View my teammates’ profiles, specifically names, profile pictures, and contact information~~  
   10. ~~Create and view a group information panel with a project name, project description, and links in the description automatically converted to anchor tags~~  
3. **Tasks**  
   1. ~~View current tasks in a simplified 3 column kanban board~~  
   2. ~~View who is assigned to each task~~  
   3. ~~Create tasks so that I can track those tasks throughout the lifetime of the group project~~  
   4. ~~Assign due dates to tasks~~  
   5. ~~Assign time estimates to tasks~~  
   6. ~~Delete tasks and mark tasks as completed~~  
   7. ~~Assign/unassign tasks so that the entire team knows who is responsible for which tasks~~  
   8. ~~Track time spent on tasks so that the team knows how much time has been spent by who on each task~~  
   9. (if time allows) View a personal window with all tasks I am assigned to across multiple groups  
4. **Stats**  
   1. ~~See a burndown chart so that I can track the team’s progress toward our goals~~  
   2. Split the burndown chart into a stacked line plot by user  
   3. Show the burndown chart by estimated time  
   4. ~~See a pie chart with how much time each team member has recorded~~  
5. **Calendar**  
   1. (if time allows) Upload my calendar from another program into Pleiades so that it can be used for finding team availabilities  
   2. ~~Manually set availabilities if I don’t use a traditional calendar~~  
   3. ~~View everyone’s availability as a calendar in a weekly view~~  
   4. ~~View due dates of tasks in the calendar view in a monthly view~~  
   5. See a color coded view of the calendar/tasks indicative of the progress made  
6. **Notifications**  
   1. (if time allows) View notifications in-app of updates and changes made to the project  
   2. (if time allows) Be able to change/customize my notification settings   
7. **Voting**  
   1. ~~Start a vote to decide project direction when consensus cannot be reached~~  
   2. Vote in a project vote  
   3. Display the results of of a vote

### Non-functional

- Architecture  
  - ~~Frontend written using the Svelte framework~~  
  - ~~Utilize components from shadcn-svelte~~  
  - ~~Backend written using Go~~  
  - ~~Use client side rendering to make the backend and frontend independent~~  
  - ~~Client-server communication will be done using websockets to allow real time updates~~  
- Appearance  
  - Use a minimalistic and modern style that is accessible and follows common, well known design patterns  
- Security  
  - ~~Use a rigorously tested library to perform password hashing and comply with security standards~~  
  - Be cognizant to avoid and mitigate buffer overflows and XSS attacks   
- Performance and Usability  
  - Animations must be smooth and we must ensure that we hit 60 fps on mobile level hardware  
  - Initial page loading times must be under 1 second including client side rendering  
  - We will perform user studies to ensure that our UI design is intuitive and understandable  
  - (if time allows) We will make our design responsive to mobile devices  
  - (if time allows) Allow using the website offline and sync when back online  
- Developer experience  
  - 60% code coverage  
- Have CI/CD so that deployments can be made quickly to the prod server  
- John will host the deployment on his personal server  
- We shall reify an impressive, useful product  
- We shall get a good grade

**1 Sprint Overview (0.5 point)**  
(a) Discuss the overview of this sprint.  
(b) Clearly state who the SCRUM master (team leader) is and briefly mention your scrum meeting schedules.  
(c) Include risks and challenges for this sprint, if you have any.  
**2 Current Sprint Detail (4.0 points)**  
(a) List all user stories to be implemented in this sprint.  
(b) For each user story, add a list of well-defined, self-contained tasks.  
(c) Ensure to include “testing” or “unit tests” task for each appropriate user story.  
(d) Add a description for each task, and clearly state which team member is assigned to the task and its workload estimation (in work hours \- make sure to distribute the total workload evenly among team members\!). Task description should be clear.  
(e) Add THREE or more detailed acceptance criteria which defines the set of conditions or statements in order for a user story to be accepted. Using “Given (some precondition) When (I do some action) Then (I expect some result).” format is strongly recommended. Typical number of successful, good acceptance criteria is about five.  
**3 Backlog (0.2 points)**  
(a) Include all the other user stories from your Product Backlog document.  
(b) (Optional, will not be graded) List tasks for each user story.  
**4 Overall Organization (0.3 points)**  
(a) Styling, clarity, right information in the right section, etc.  
(b) Please make sure to include your project name, team number and the names of all your team members.  
