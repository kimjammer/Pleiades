**Team 5 Sprint 3 Planning**  
**Pleiades**  
Cate Harrison, John Kim, Henry Rovnyak, Ethan Dawes

## Sprint Overview

In our final sprint, we will be finishing up key core functions such as polls, personal task management across projects, and burndown charts. We will also be adding functionality for features that will make the app more usable, such as notifications and password recovery. These items will round out the capabilities of Pleiades, making it a fully functional platform for effective communication, teamwork, and personal time management.

**Scrum Master:** Henry Rovnyak

**Meeting Plan:** Monday 11:30 AM, Wednesday 11:30 AM

### Risks and Challenges

This sprint will involve finishing and polishing what has been completed in our first and second sprints. Luckily, since much of our core functionality was completed last sprint, we will be able to reuse many of our components to swiftly complete user stories. One thing to watch out for is that since this is our final sprint, we must make sure to fully complete our user stories and to perform thorough testing to find and fix all of the bugs that may exist in our system.

### 

### User Story \#1.4

As a user, I want to be able to reset my password so that I can recover my account

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Setup third party email service | 1 hr | John |
| 2 | Implement backend logic to generate password recovery links | 2 hr | John |
| 3 | Implement password change API | 2 hr | John |
| 4 | Create frontend page to send recovery link | 2 hr | John |
| 5 | Create frontend page to enter new password | 2 hr | John |
| 6 | Write tests for password recovery | 1 hr | John |

Acceptance Criteria:

1. Given that you are on the login page, you can navigate to the request password reset page.  
2. Given that you enter an email associated with an account, an email will be sent to the address with a special link to change your password.  
3. Given that you open the special link, you will be able to enter a new password.  
4. Given that the password has been changed, you can no longer login with the old password but you can login with the new password.

### User Story \#4.2

As a user, I want to be able to split the burndown chart into a stacked line plot by user

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Refactor due dates to have a time, instead of just a day | 3 hr | John |
| 2 | Refactor session aggregation logic to support iterating over users | 3 hr | John |
| 3 | Add toggling to stacked line plot by user | 2 hr | John |
| 4 | Write tests for burndown chart | 1 hr | John |

Acceptance Criteria:

1. Given there is at least one task with both a time estimate and a due date, you can toggle the burndown chart to show a stacked line plot by user instead of team total time.  
2. Given the burndown chart is toggled to stacked mode, each user’s lines show as stacked and the topmost line is the same as in total mode.  
3. Given the burndown chart is toggled into stacked mode, it also updates in real time.  
4. Given the burndown chart is toggled into stacked mode, legends and tooltips display which user each line corresponds to.

### User Story \#4.3

As a user, I want to be able to show the burndown chart by estimated time

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Refactor session aggregation logic to support iterating over tasks | 3 hr | John |
| 2 | Add toggling to stacked line plot by task | 2 hr | John |
| 3 | Ensure that timestamps are handled consistently across all components | 2 hr | John |
| 4 | Write end to end tests for chart | 2 hr | John |

Acceptance Criteria:

1. Given there is at least one task with both a time estimate and a due date, you can toggle the burndown chart to show a stacked line plot by task instead of all tasks.  
2. Given the burndown chart is toggled into task mode, each task shows its own line stacked on top of each other.  
3. Given the burndown chart is toggled into task mode, it also updates in real time.  
4. Given the burndown chart is toggled into task mode, legends and tooltips display which task each line corresponds to.

### User Story \#7.2

As a user, I want to vote in a project vote

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Set up the UI for project votes | 3 hr | Henry |
| 2 | Make voting update the server | 2 hr | Henry |
| 3 | Make the UI responsive to votes being selected | 2 hr | Henry |

Acceptance Criteria:

1. For each poll option, the poll option should be visible on screen with options to select “approve”, “neutral”, and “disapprove” (may be icons)  
2. Clicking on a poll option should highlight that it was selected  
3. Clicking on a different poll option will deselect the previous poll option and select the new one

### User Story \#7.3

As a user, I want to display the results of a poll

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Allow changing the end date of a poll (for testing) | 1 hr | Henry |
| 1 | Disallow voting after the given end date | 1 hr | Henry |
| 2 | After the end date, the poll options should be sorted based on approval | 2 hr | Henry |
| 3 | Display a stacked bar chart for each option displaying approvals | 4 hr | Henry |

Acceptance Criteria:

1. Before the due date, the original voting UI is displayed  
2. After the due date, the new results UI is displayed  
3. The results should be sorted based on approval  
4. Each poll option should display a stacked bar chart visualizing how many people selected “approve” vs “neutral” vs “disapprove”

### User Story \#6.1

As a user, I want to view notifications in-app of updates and changes made to the project

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Update the Websockets backend as well as frontend interface to allow broadcasting notifications from both the frontend and backend | 5 hr | Henry |
| 2 | Set up the UI for displaying notifications | 3 hr | Henry |
| 3 | Hook into various project events and make them display notifications | 5 hr | Henry |

Acceptance Criteria:

1. A user joining a project should display a notification to other users  
2. A poll ending should display a notification to other users  
3. Assigning a task to a user should display a notification to that user

### User Story \#3.9

As a user, I want to be able to view a personal window with all tasks I am assigned to across multiple groups

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Create personal window UI | 4 hr | Cate |
| 2 | Write tests for personal window | 2 hr | Cate |
| 3 | Implement view filtering for different projects | 2 hr | Cate |
| 4 | Implement filtering for assigned tasks versus all tasks in a project. | 3 hr | Cate |

Acceptance Criteria:

1. Given that a user is assigned to tasks in different projects, all tasks will appear in their personal window.  
2. Given that a user has tasks in multiple projects, they can filter which projects to see.  
3. Given the user has assigned tasks, they can modify the tasks in their personal window the same as in the project window.  
4. Given that a user hovers their mouse over a task, a tooltip will appear with further details about the task.

### User Story \#5.5

As a user, I want to see a color coded view of the calendar/tasks indicative of the progress made

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Implement color changing of project names based on progress | 4 hr | Cate |
| 2 | Implement progress bar color changing in task cards based on session completion | 4 hr | Cate |
| 3 | Write tests for color changing UI |  2 hr | Cate |

Acceptance Criteria:

1. Given that a user has not begun working on a task, the default color will be black.  
2. Given that progress has been made on the task, a logical color scheme will be applied according to progress.  
3. Given that a task is moved to the “Completed” section, it will automatically turn green regardless of the progress bar.

### User Story \#6.2

As a user, I want to be able to change/customize my notification settings.

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Create UI for settings options | 4hr | Cate |
| 2 | Implement filtering for notifications | 4hr | Cate |
| 3 | Write tests for notification settings | 2hr | Cate |

Acceptance Criteria:

1. Given that the user is in a project, they can turn off/on notifications for when other users join the project.  
2. Given that a user is in a project, they can turn off/on notifications for polls ending soon.  
3. Given that a user is in a project, they can turn off/on notifications for being assigned to tasks.

### User Story \#3.1

As a user, I want to be able to login with social accounts (Google)

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Configure Google Dev Console | 1 hr | Ethan |
| 2 | Create oauth callback endpoint to login | 1.5 hr | Ethan |
| 3 | Create frontend for login/account creation | 2 hr | Ethan |
| 4 | Set profile pic from google account | 0.5 hr | Ethan |
| 5 | Edge case handling (expiry,denial, etc) | 1 hr | Ethan |
| 6 | Automated testing | 1.5 hr | Ethan |
| 7 | Create oauth callback endpoint to register account with google | 1.5 hr | Ethan |

9 hr  
Acceptance Criteria:

1. Given the user clicks “log in with Google” and they have a google-linked account, they will be logged in as usual  
2. Given the user clicks “register with Google” an account will be created with their Google name, email, and profile picture  
3. Given the user initiate a login flow but denies permission, the user will see an error indicating so

### User Story \#2.3

As a user, I want to be able to invite people directly in the app using their email

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Modify the UI to comprehensively combine the teammate list, invite button and invite by email ability | 1 hr | Ethan |
| 2 | Integrate with Purdue directory to get email from name (make endpoint to proxy & parse HTML from Purdue) | 1 hr | Ethan |
| 3 | Send an email to that person with the join link | 1 hr | Ethan |
| 4 | Write tests for 1 & 2 | 1.5 hr | Ethan |

4.5 hr  
Acceptance Criteria:

1. Given I enter an email to invite, that recipient will get an email with an invite code  
2. Given I enter the name of a student at Purdue, a drop down appears with possible people matches.  
3. Given I click on a suggested name, an email will be sent to their email retrieved from the Purdue directory.

### User Story \#5.1

As a user, I want to be able to upload my calendar from another program into Pleiades so that it can be used for finding team availabilities

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Fix calendar storage format to eliminate need to convert between client and server data format (refactor generally) | 3 hr | Ethan |
| 2 | Fix bug where monday times cannot be unselected | 2 hr | Ethan |
| 3 | Add button to UI | 0.5 hr | Ethan |
| 4 | Package input selector into NPM package to help open source community | 2.5 hr | Ethan |
| 5 | Write test mocks for importing from google calendar | 2 hr | Ethan |
| 6 | Convert gcal api to compatible data format | 2 hr | Ethan |

12 hr  
Acceptance Criteria:

1. Given the user is on the availability tab and clicks “import from Google” or similar and their account is not connected to Google, they will be given an auth confirmation screen  
2. Given the user is on the availability tab and clicks “import from Google” or similar and their account is connected to Google, their availability should be overwritten with “free” where there are no calendar events and “busy” where there is a calendar event  
3. Given the user has calendar events marked as “free”, those should not be marked as “unavailable” in the availability calendar

### User Story \#8.1

As a user, I want to be able to see how many people have created an account, visited the site, and hours logged to projects to provide credibility and provide feedback

| \# | Description | Estimated Time | Owner(s) |
| :---- | :---- | :---- | :---- |
| 1 | Implement a stats retrieval endpoint | 1 hr | Ethan |
| 2 | Create an analytics dashboard for various interactions of interest | 2 hr | Ethan |
| 3 | Conduct user studies using the stats collection | 2 hr | Ethan |
| 4 | Create tests for stats | 1 hr | Ethan |

6 hr  
Acceptance Criteria:

1. Given I visit the home page, I should see stats for how many people have created an account, visited the site, and hours logged  
2. Given I perform one of the above actions, I should see the homepage updated to reflect the latest values on the next reload  
3. Given the nav header is visible, I should see an additional button linking to a Google Form to provide feedback

### Remaining Backlog

### Functional

As a user, I want to be able to

1. **Account**  
   1. ~~Create an account with a email, real name, and password~~  
   2. ~~Log in with a email and password so that I can see information for me and my team~~  
   3. ~~(If time allows) login with social accounts (Google)~~  
   4. ~~(If time allows) Reset my password so that I can recover my account~~  
   5. ~~Upload a profile picture~~  
2. **Group**  
   1. ~~Create a group so that I can begin my project~~  
   2. ~~Generate an invite link to the group so that I can collaborate with other students~~  
   3. ~~(If time allows) Invite people directly in the app using their email~~  
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
   9. ~~(if time allows) View a personal window with all tasks I am assigned to across multiple groups~~  
4. **Stats**  
   1. ~~See a burndown chart so that I can track the team’s progress toward our goals~~  
   2. ~~Split the burndown chart into a stacked line plot by user~~  
   3. ~~Show the burndown chart by estimated time~~  
   4. ~~See a pie chart with how much time each team member has recorded~~  
5. **~~Calendar~~**   
   1. ~~(if time allows) Upload my calendar from another program into Pleiades so that it can be used for finding team availabilities~~  
   2. ~~Manually set availabilities if I don’t use a traditional calendar~~  
   3. ~~View everyone’s availability as a calendar in a weekly view~~  
   4. ~~View due dates of tasks in the calendar view in a monthly view~~  
   5. ~~See a color coded view of the calendar/tasks indicative of the progress made~~  
6. **~~Notifications~~**  
   1. ~~(if time allows) View notifications in-app of updates and changes made to the project~~  
   2. ~~(if time allows) Be able to change/customize my notification settings~~   
7. **~~Voting~~**  
   1. ~~Start a vote to decide project direction when consensus cannot be reached~~  
   2. ~~Vote in a project vote~~  
   3. ~~Display the results of of a vote~~  
8. **~~Analytics~~**  
   1. ~~As a user, I want to be able to see how many people have created an account, visited the site, and hours logged to projects to provide credibility~~

### Non-functional

- Architecture  
  - ~~Frontend written using the Svelte framework~~  
  - ~~Utilize components from shadcn-svelte~~  
  - ~~Backend written using Go~~  
  - ~~Use client side rendering to make the backend and frontend independent~~  
  - ~~Client-server communication will be done using websockets to allow real time updates~~  
- Appearance  
  - ~~Use a minimalistic and modern style that is accessible and follows common, well known design patterns~~  
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
  - ~~60% code coverage~~  
- Have CI/CD so that deployments can be made quickly to the prod server  
- ~~John will host the deployment on his personal server~~  
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

