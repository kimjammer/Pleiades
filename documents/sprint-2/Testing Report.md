**Team 5 Sprint 2 Testing Report**  
**Pleiades**  
Cate Harrison, John Kim, Henry Rovnyak, Ethan Dawes  
Setup:

1. cd frontend  
2. npm i  
3. npm run dev  
4. (in another session) cd backend  
5. go run \-tags TEST .  
6. Open localhost:5173/BASE\_PATH in browser

Automated Tests:

1. cd backend  
2. go test .  
3. cd ../frontend  
4. npm run build  
5. npm run test

### User Story \#1.5

As a user, I want to be able to upload a profile picture.

Acceptance Criteria:

1. ### Given that the user uploads a valid image, it will be displayed as the profile picture

2. ### Given that the user tries to upload an invalid file, it will not work.

3. ### Given that a user has not uploaded a profile picture, they will have a default image instead.

#### Test Case

This test case tests acceptance criteria 1, given that the user uploads a valid image, it will be displayed as the profile picture, and 2, given that the user tries to upload an invalid file, it will not work, and 3, given that a user has not uploaded a profile picture, they will have a default image instead.

#### Input

Open the website, and log in to a preexisting account WITHOUT a profile picture. Click on the icon for the account page, which is the current user avatar. Click on “Upload profile picture.” Select an image file to upload as your profile picture and submit. Reload the page.

#### Output

Upon first logging in, the user avatar will be a set of initials.  
When a new picture is uploaded, only valid image files pop up even in folders with different file types.  
The new profile picture will now appear in the UserAvatar component.

#### Acceptance

Pass.

### User Story \#3.1

As a user, I want to view current tasks in a simplified 3 column kanban board

Acceptance Criteria:

1. The project page will show the three column kanban board  
2. Each task will display as a card in the correct column  
3. Multiple connected users will display the same information

#### Test Case

This test case tests all three acceptance criteria.

#### Input

Open an empty project as a logged in user. Create a new task with arbitrary details. Click the task. Duplicate the tab and click the task in the new tab.

#### Output

Upon opening the empty project, the user should see the three column kanban board. Upon creating a new task, the task should appear under the “Not started” column. Clicking the task will display the same information as was given during creation of the task. Duplicating the tab and clicking the task should show the same information as on the previous tab.

#### Acceptance

Pass.

### User Story \#3.2

As a user, I want to view who is assigned to each task

Acceptance Criteria:

1. If a task is not assigned to a user, the card should not display a user  
2. If a task is assigned a user, the card should display it in a line  
3. The user should be displayed as their full real name

#### Test Case

This test case tests all three acceptance criteria.

#### Input

Open an empty project as a logged in user. Create a new task with arbitrary details, but without assigning a user. Click the task and click “Join Task”.

#### Output

Upon creating a new task and clicking the task, the task should not display any assigned user. Clicking “Join Task” will make the current user’s first name and last name appear in the task. 

#### Acceptance

Pass.

### User Story \#3.6

As a user, I want to delete tasks and mark tasks as completed

Acceptance Criteria:

1. Dragging and dropping cards will move the card to the column it was dropped in  
2. For a different user, the card will be moved to that column on their screen  
3. Right clicking a card will show a custom tooltip  
4. Clicking delete on the tooltip will delete the task  
5. The task will appear deleted on the screen of another user

#### Test Case

This test case tests all five acceptance criteria.

#### Input

Open an empty project as a logged in user. Create a new task with arbitrary details. Duplicate the tab. In the new tap, drag and drop the task to the “Done” column. Switch to the previous tab. Right click on the task. Click “Delete”. Switch to the duplicated tab.

#### Output

Upon dragging and dropping the task, the task should move to the “Done” column. Upon switching tabs, the task should appear in the “Done” column in that tab too. Upon right clicking on the task, a tooltip with the button “Delete” should appear. Upon clicking the button, the task will be deleted in both tabs.

#### Acceptance

Pass.

### User Story \#3.7

As a user, I want to be able to assign/unassign tasks so that the entire team knows who is responsible for which tasks.

Acceptance Criteria:

1. ### Given that a user is not assigned to a task, they may join a task.

2. ### Given that a user is assigned to a task, they may leave a task.

3. ### Given that a user is already assigned to a task, they may not join it twice.

4. ### Given that there are users assigned to a task, all users in the workspace can see the current assigned users.

#### Test Case

This test case tests acceptance criteria 2, given that a user is assigned to a task, they may leave a task.

#### Input

Open the website, and log in to a preexisting account and navigate to a project page. Create a new task, and assign the current user to that task.  
Open the task card, and click “Leave Task” at the bottom. 

#### Output

The current user’s name will disappear from the task.

#### Acceptance

Pass.

#### Test Case

This test case tests acceptance criteria 1, given that a user is not assigned to a task, they may join a task, and 3, given that a user is already assigned to a task, they may not join it twice.

#### Input

Open the website, and log in to a preexisting account and navigate to a project page. Create a new task, and do NOT assign the current user to that task.  
Open the task card, and click “Join Task” at the bottom. 

#### Output

The current user’s name will be added to the task. The “Join Task” button also turns into “Leave Task,” preventing a user from joining more than one task.

#### Acceptance

Pass.

#### Test Case

This test case tests acceptance criteria 4, given that there are users assigned to a task, all users in the workspace can see the current assigned users.

#### Input

Open the website, and log in to a preexisting account and navigate to a project page with at least two members. Create a new task, and assign 2+ members to that task.  
Log out and log back in as a different user in the same project. Open the project and open the previously created task card.

#### Output

The second user will see the same task card as the first user, displaying all current members of the task.

#### Acceptance

Pass.

### User Story \#7.1

As a user, I want to be able to start a vote to decide project direction when consensus cannot be reached.

Acceptance Criteria:

1. Given that a user is in a project, they can create a poll.  
2. Given that a poll exists in a project, users in that project can see that it exists.  
3. Given that a poll does not have at least a title and two options, it cannot be created.  
4. Given that a poll is created, it has a time at which it will end.

#### Test Case

This test case tests acceptance criteria 3, given that a poll does not have at least a title and two options, it cannot be created, and 4, given that a poll is created, it has a time at which it will end.

#### Input

Open the website, and log in to a preexisting account and navigate to a project page, and click on the Polls tab.  
Click “Create a new Poll” and immediately select “submit”

#### Output

Red text will appear telling the user to input a title and at least two options and a due date.

#### Acceptance

Pass.

#### Test Case

This test case tests acceptance criteria 1, given that a user is in a project, they can create a poll, and 2, given that a poll exists in a project, users in that project can see that it exists.

#### Input

Open the website, and log in to a preexisting account and navigate to a project page, and click on the Polls tab.  
Click “Create a new Poll” and input information for the title, due date, and options. The options field must have at least one comma separating two options.  
Click “Submit”

#### Output

A new button appears on the Polls tab indicating the existence of a new poll.

#### Acceptance

Pass.

### User Story \#3.4

As a user, I want to assign due dates to tasks

Acceptance Criteria:

1. Given that a task exists and has no due date, there is a button to input a due date.  
2. Given that a task exists and has a due date, there is a chip displaying the due date.  
3. Given that a task exists and has a due date, it can be changed by clicking on it.  
4. Given that a task has a due date, it is displayed in your local timezone.

#### Test Case

This test case tests acceptance criteria 1, that you will be able to input a new due date

#### Input

Open the website, log in, and open a project.  
Create a new task, without specifying the due date.  
Click on the newly created task to expand it, and observe the due date component.

#### Output

The due date chip says “Due Date” with a plus icon.

#### Acceptance

Pass.

#### Test Case

This test case tests acceptance criteria 2, 3, and 4, that you can see the due date in your local timezone and are able to change it

#### Input

Open the website, log in, and open a project.  
Create a new task, and specify a due date of today.  
Click on the newly created task to expand it, and observe the due date component.  
Click on the due date component and choose a new due date.

#### Output

The due date chip displays the month and day of today, even if UTC is on another day.  
Clicking the due date component displays a calendar and when a new due date is chosen, the component updates right away.

#### Acceptance

Pass.

### User Story \#3.5

As a user, I want to be able to assign time estimates to tasks

Acceptance Criteria:

1. Given that a task exists and has no time estimate, there is a button to add a time estimate.  
2. Given that a task exists and has a time estimate, there is a chip displaying the time estimate.  
3. Given that a task exists and has a time estimate, it can be changed by clicking on it.

#### Test Case

This test case tests acceptance criteria 1, that you can add a new time estimate

#### Input

Open the website, log in, and open a project.  
Create a new task, without specifying the time estimate.  
Click on the newly created task to expand it, and observe the due date component.

#### Output

The due date chip says “Time Estimate” with a plus icon.

#### Acceptance

Pass.

#### Test Case

This test case tests acceptance criteria 2 and 3, that you can see the current time estimate and that you can change it

#### Input

Open the website, log in, and open a project.  
Create a new task, and specify a time estimate.  
Click on the newly created task to expand it, and observe the due date component.  
Click on it and change the value to a new value.

#### Output

The due date chip says the current time estimate.  
When the chip is clicked, it provides an input to type into, and when a new value is entered, the chip updates right away.

#### Acceptance

Pass.

### User Story \#3.8

As a user, I want to be able to track time spent on tasks so that the team knows how much time has been spent by who on each task

Acceptance Criteria:

1. Given a task, you can start a timer to record your work session.  
2. Given a timer is running, you can stop the timer to complete your work session.  
3. Given at least 1 user has at least 1 session for a task, you can see the total amount of time spent by each person on that task.  
4. Given that at least 1 session exists for a task, you can see the progress to the time estimate.

#### Test Case

This test case tests acceptance criteria 1 and 2, that you can start and stop sessions

#### Input

Open the website, log in, and open a project.  
Create a new task.  
Click the task to expand it, and press the play button to start recording a session.  
Some time later, press it again to stop the session.

#### Output

When the play button is pressed, it will display a timer that increments every second.  
When the stop button is pressed, the timer will go away.

#### Acceptance

Pass.

#### Test Case

This test case tests acceptance criteria 3, that you can see worked sessions

#### Input

Open the website, log in, and open a project.  
Create a new task.  
Click the task to expand it, and press the play button to start recording a session.  
Some time later, press it again to stop the session.  
Hover over the Session text

#### Output

A popup will appear displaying the session that was just recorded, with your user profile and the length of the session

#### Acceptance

Pass.

#### Test Case

This test case tests acceptance criteria 4, that you can see the progress to the time estimate

#### Input

Open the website, log in, and open a project.  
Create a new task with a time estimate of 0.1 hours.  
Click the task to expand it, and press the play button to start recording a session.  
Some time later, press it again to stop the session.

#### Output

When the session is stopped, a progress bar appears at the bottom of the task card that corresponds to how much time was spent towards the time estimate.

#### Acceptance

Pass.

### User Story \#4.1

As a user, I want to be able to see a burndown chart so that I can track the team’s progress toward our goals.

Acceptance Criteria:

1. Given at least 1 task with both a time estimate and a session, you can see a burndown chart of total time estimate vs time spent on a daily interval.  
2. Given there is no task with both a time estimate and a session, the graph component shows a helpful “N/A” message to the user.  
3. Given a new time estimate or session is recorded, the graph updates in real time.

#### Test Case

This test case tests acceptance criteria 2, that the graph will correctly not display if there is no data to display.

#### Input

Open the website, log in, and open a project with no tasks.  
Switch to the Stats tab

#### Output

The burndown graph component says “Create a task with a due date and time estimate to see the burndown chart.”

#### Acceptance

Pass.

#### Test Case

This test case tests acceptance criteria 1 and 3, that the graph will show a chart of time estimate vs time spent, updating in real time.

#### Input

Open the website, log in, and open a project.  
Create a new project with a due date 5 days in the future, and a time estimate of 1 hours.  
Expand the task, and press start to begin a session.  
Some time later, press stop to end the session.  
Switch to the Stats tab, without reloading the page.

#### Output

The x axis of the graph will span from the day the project was created to 5 days from today. The y axis of the graph will span from 0 to 1 hours. The blue “ideal” line increases steadily from the creation of the project to the day before its due date, then it reaches the time estimate on the due date. The red “actual” line will be 0 until the next day, where it will reach the recorded time session and stay flat.

#### Acceptance

Pass.

### User Story \#2.5

As a user, I want to be able to be directed to the login page when clicking an invite link if not signed in and automatically being added to the group after signing in

Acceptance Criteria:

1. Given the user visits the /join page while signed out, they are not immediately directed to the landing page (as implemented last sprint). Instead, they see the join confirmation page. (manual)  
2. Given the user clicks the accept button, they are redirected to the login page (manual)  
3. Given the user did steps 1 & 2 of the acceptance criteria, after the user signs in, they are taken directly to the project page for the project they just accepted & become a project member. (manual)

#### Test Case

This test case tests all 3 acceptance criteria

#### Input

1. Log in, open or create a project  
2. Go to “settings” tab and make an invite link  
3. Copy the link and close the modal  
4. Sign out  
5. Paste the link in the address bar

#### Output

1. Observe the confirmation page and confirm that the title and description matches the project you were just in (criteria 1\)  
2. Click the “join” button  
3. Confirm that the address bar reads `/registration` (criteria 2\)  
4. Register an account  
5. Observe you are now on the project page for the project you joined (criteria 3\)

#### Acceptance

Pass.

### User Story \#2.9

As a user, I want to be able to view my teammates’ profiles, specifically names, profile pictures, and contact information

Acceptance Criteria:

1. Given I am looking at my teammate’s info tab, if I click the “contact” button, my email client opens to send them mail (manual)  
2. Given I am in the project, I will see a tab where the group info is displayed. (manual)  
3. Given I am looking at the teammate info tab, I will see each user’s names, profile pictures, and a method to contact them is displayed (manual)

#### Test Case

This test case tests all 3 acceptance criteria

#### Input

1. Log in, open or create a project  
2. See previously outlined procedure for setting profile picture and adding another user to project

#### Output

1. Go to “settings” tab (criteria 2\)  
2. Observe yourself in the list of project members & your profile picture or initials (criteria 3\)  
3. Click “contact” beneath your name and wait for email client to open (criteria 1\)

#### Acceptance

Pass.

### User Story \#3.3

As a user, I want to be able to create tasks so that I can track those tasks throughout the lifetime of the group project

Acceptance Criteria:

1. Given the user is on the tasks tab of the project page, they will see a button to create a new task  
2. Given the user presses this button to create a new task, they are presented with a modal with fields such as title, description, due date, time estimate, and assignees.  
3. Given the user is looking at the task creation modal, they will not be able to create a task and will be presented with an error if they leave the title blank  
4. Given the user has typed a duration (5 hours/5hr) or a due date (next tuesday/Jan 15\) into the title field, the respective fields later in the form will be populated with this information

#### Test Case

This test case tests all criteria

#### Input

1. Log in, open or create a project. Stay on default “tasks” tabs

#### Output

1. Find & click the “Create a new task” button (criteria 1\)  
2. Observe the modal has all the required fields (criteria 2\)  
3. Click the “create” button and observe the helpful errors that alert you that the title must be filled (criteria 3\)  
4. Type “analyse stats next tuesday 1 hour 30m” and wait a second. Confirm that the “due date” field has been filled to next tuesday, estimated time is 1.5, and title is simply “analyse stats” (criteria 4\)

#### Acceptance

Pass.

### User Story \#4.4

As a user, I want to be able to see a pie chart with how much time each team member has recorded

Acceptance Criteria:

1. Given I am looking at the designated tab on the project page and members have recorded time, there will be a pie chart, where each pie slice represents the cumulative time a team member spent on any task in this project (manual)  
2. Given I hover over a slice on the pie chart, a tooltip appears with the user's name and how much time they spent (manual)  
3. Given I am looking at the designated tab on the project page and no members have recorded time, there will be a message indicating so and encouraging members to track time (manual)

#### Test Case

This test case tests criteria 3

#### Input

1. Create a new project & open it  
2. Go to the “stats” tab

#### Output

Observe the notice encouraging you to record time

#### Acceptance

Pass.

#### Test Case

This test case tests criteria 1 & 2

#### Input

1. Open a project with tasks  
2. Log time for those tasks  
3. Switch to another account and record more time  
4. Navigate to the “stats” tab

#### Output

Observe the pie chart of your cumulative time (criteria 1). Hover over a slice and observe your name and time worked on the project in hours (criteria 2\)

#### Acceptance

Pass.

### User Story \#5.4

As a user, I want to be able to view due dates of tasks in the calendar view in a monthly view

Acceptance Criteria:

1. Given I am looking at the designated tab on the project page, I will see a month-view calendar  
2. Given users have created tasks with due dates, each of those incomplete tasks titles is placed on its respective date.  
3. Given I press the left and right buttons at the top of the page, the month I am looking at changes

#### Test Case

This test case tests all criteria

#### Input

1. Have a project with multiple tasks, between multiple months  
2. Switch to the “calendar” tab

#### Output

1. Confirm that the dates match up with the current month (criteria 1\)  
2. Confirm that the task due dates match up with their respective date on the calendar (criteria 2\)  
3. Hover over the month input number, and click the left or right button to go to the next/previous month and observe the tasks that are due in that month (criteria 3\)

#### Acceptance

Pass.

**1 Version Control (1.0 point)**  
(a) All team members are required to consistently use version control system.  
**2 Weekly Reports (1.5 points)**  
(a) Your current scrum master should submit weekly team reports to your project coordinator and  
the head TA.  
(b) Each individual should submit weekly individual reports to their project coordinator.  
(c) Make sure to tag each task with your version control commit identifiers (e.g. git commit hash) as necessary.  
(d) Both reports should be submitted via email by 11:58pm EVERY Sunday and no points will be given for late submissions.  
**3 Implementation and Testing (9.0 points)**  
(a) Your implementation and demo should be based on the Sprint Planning Document. Please implement the tasks that you have planned in the document.  
(b) Each team member must present and explain the tasks they have worked on during the sprint.  
(c) If you have any automated tests, then you should demonstrate them during the meeting.  
(d) For any manual tests, you should submit a report of the test cases that were executed during the sprint. List a user story and its test cases. Each test case should include a short description, its input, observed output and should indicate if it was an acceptable output (yes or no).  
(e) The manual testing report should be submitted anytime before the meeting.  
**4 Overall Presentation (1.5 points)**  
(a) Ensure to prepare your demo before the meeting and do a dry-run.  
**Notes**  
(a) Please arrange to meet with your assigned grader on the suggested meeting day (usually Friday) or one day before it.  
(b) If you have some tasks based on machine learning or data mining algorithms, and cannot avoid certain error rates, it is understandable for your application not to always produce correct  
results. No points will be deducted for the tasks which work reasonably well.  
