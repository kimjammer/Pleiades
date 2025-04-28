**Team 5 Sprint 3 Testing Report**  
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

### User Story \#TEMPLATE

As a user,

Acceptance Criteria:

1. ### Given that 

#### Test Case

#### Input

#### Output

#### Acceptance

Pass.

### User Story \#1.4

As a user, I want to be able to reset my password so that I can recover my account.

Acceptance Criteria:

1. Given that you are on the login page, you can navigate to the request password reset page.  
2. Given that you enter an email associated with an account, an email will be sent to the address with a special link to change your password.  
3. Given that you open the special link, you will be able to enter a new password.  
4. Given that the password has been changed, you can no longer login with the old password but you can login with the new password.

#### Test Case

This test case confirms that you will receive a password reset email, and that it changes your password as expected.

#### Input

Create an account, log out, and go to the login screen. Click on “Forgot your Password?” and enter the email associated with the account. Click send email. Follow the link in the email and enter a new password. Attempt to login with the old password and the new password.

#### Output

The old password fails the login. The new password will pass the login.

#### Acceptance

Pass.

### User Story \#4.2

As a user, I want to be able to split the burndown chart into a stacked line plot by user.

Acceptance Criteria:

1. Given there is at least one task with both a time estimate and a due date, you can toggle the burndown chart to show a stacked line plot by user instead of team total time.  
2. Given the burndown chart is toggled to stacked mode, each user’s lines show as stacked and the topmost line is the same as in total mode.  
3. Given the burndown chart is toggled into stacked mode, it also updates in real time.  
4. Given the burndown chart is toggled into stacked mode, legends and tooltips display which user each line corresponds to.

#### Test Case

This test case tests that the toggle is functional and that the graph will change modes accordingly.

#### Input

In a project, create a task with a due date and an estimated time. On two different accounts, both record a session for that project. In the Stats tab, click the “Breakout Users” toggle.

#### Output

The “actual” line of the burndown chart will turn into a stacked line chart where the top most line is still consistent with the “actual” line, but the session times of each user are broken out.

#### Acceptance

Pass.

### User Story \#4.3

As a user, I want to be able to show the burndown chart by estimated time.

Acceptance Criteria:

1. Given there is at least one task with both a time estimate and a due date, you can toggle the burndown chart to show a stacked line plot by task instead of all tasks.  
2. Given the burndown chart is toggled into task mode, each task shows its own line stacked on top of each other.  
3. Given the burndown chart is toggled into task mode, it also updates in real time.  
4. Given the burndown chart is toggled into task mode, legends and tooltips display which task each line corresponds to.

#### Test Case

This test case tests that the toggle is functional and that the graph will change modes accordingly.

#### Input

In a project, create two tasks with different due dates and estimated times. In the Stats tab, click the “Breakout Tasks” toggle.

#### Output

The “ideal” line of the burndown chart will turn into a stacked line chart where the top most line is still consistent with the “ideal” line, but the estimated times of each task are broken out.

#### Acceptance

Pass.

### User Story \#7.2

As a user, I want to vote in a project vote

Acceptance Criteria:

1. For each poll option, the poll option should be visible on screen with options to select “approve”, “neutral”, and “disapprove” (may be icons)  
2. Clicking on a poll option should highlight that it was selected  
3. Clicking on a different poll option will deselect the previous poll option and select the new one

#### Test Case

This test case tests all three acceptance criteria.

#### Input

Create a new poll with three options. Vote approve, neutral, and disapprove for the options in that order. Change the votes to neutral, disapprove, and approve in that order.

#### Output

Creating the poll will show three buttons for each poll option. Clicking on the buttons to vote will highlight them as green, yellow, or red for approve, neutral, and disapprove respectively. Changing the vote to a different option will select that option instead.

#### Acceptance

Pass.

### User Story \#7.3

As a user, I want to display the results of a poll

Acceptance Criteria:

1. Before the due date, the original voting UI is displayed  
2. After the due date, the new results UI is displayed  
3. The results should be sorted based on approval  
4. Each poll option should display a stacked bar chart visualizing how many people selected “approve” vs “neutral” vs “disapprove”

#### Test Case

This test case tests all four acceptance criteria.

#### Input

Create a new poll with three options. Vote disapprove, neutral, and approve for the options in that order. Log out and log in as another user. Vote neutral, approve, and disapprove for the options in that order. Right click on the poll and select “End poll”.

#### Output

Before the poll is over, the voting UI will be displayed. After selecting “End poll”, users will no longer be able to vote in the poll and the poll will display the results. The results will be sorted in order of second option, then third option, then first option. Each poll option will display a stacked bar chart showing how people ranked each option.

#### Acceptance

Pass.

### User Story \#6.1

As a user, I want to view notifications in-app of updates and changes made to the project

Acceptance Criteria:

1. A user joining a project should display a notification to other users  
2. A poll ending should display a notification to other users  
3. Assigning a task to a user should display a notification to that user

#### Test Case

This test case tests all three acceptance criteria.

#### Input

Log in as a user and verify that all notifications are enabled in the options. Create an invite link to a project. In a separate window, log in as another user, open the invite link, and join the project. As the new user, create a poll, right click, and end the poll. As the original user, create a new task while assigning it to the new user.

#### Output

When the new user joins the project, the original user will see a notification that they joined the project. When the new user ends the poll, both users will see a notification that the poll ended. When the original user assigns a task to the new user, the new user will see a notification that they were assigned the task.

#### Acceptance

Pass.

### User Story \#6.2

As a user, I want to be able to change/customize my notification settings.

Acceptance Criteria:

1. Given that the user is in a project, they can turn off/on notifications for when other users join the project.  
2. Given that a user is in a project, they can turn off/on notifications for polls ending soon.  
3. Given that a user is in a project, they can turn off/on notifications for being assigned to tasks.

#### Test Case

This test case will test all 3 acceptance criteria.

#### Input

Log in as a user, navigate to the account page, and turn off all 3 notification settings. Create an invite link to a project. In a separate window, log in as another user, open the invite link, and join the project. As the new user, create a poll, right click, and end the poll. As the original user, create a new task while assigning it to the new user.

#### Output

There will be no pop-up notifications

#### Acceptance

Pass.

### User Story \#3.9

As a user, I want to be able to view a personal window with all tasks I am assigned to across multiple groups

Acceptance Criteria:

1. Given that a user is assigned to tasks in different projects, all tasks will appear in their personal window.  
2. Given that a user has tasks in multiple projects, they can filter which projects to see.  
3. ~~Given the user has assigned tasks, they can modify the tasks in their personal window the same as in the project window.~~ Given the user has assigned tasks, they can click on a task to navigate to the task window in that project.  
4. Given that a user hovers their mouse over a task, a tooltip will appear with further details about the task.

#### Test Case

This test case tests all four acceptance criteria

#### Input

Log in as a user and ensure the user has tasks assigned in at least 2 different projects. Navigate the account page and scroll down to the personal window. Observe that all assigned tasks appear on the calendar.  
Click on “Filter” and toggle one or more projects to off. Click “Apply Filter” and observe that the appropriate tasks have disappeared.  
Hover the mouse over any task and observe the task details  
Click on a task and observe that the user is directed to the project’s task page.

#### Output

The user will be able to filter projects in the personal window, view hover card information about the task, and navigate to the task’s project via clicking.

#### Acceptance

Pass.

### User Story \#5.5

As a user, I want to see a color coded view of the calendar/tasks indicative of the progress made

Acceptance Criteria:

1. Given that a user has not begun working on a task, the default color will be black.  
2. Given that progress has been made on the task, a logical color scheme will be applied according to progress.  
3. Given that a task is moved to the “Completed” section, it will automatically turn green regardless of the progress bar.

#### Test Case

This test case tests all three acceptance criteria

#### Input

Log in as a user and navigate to a project page. Create a new project. Click start session. Wait a few minutes. Stop the session. Move the task through all three kanban Columns.

#### Output

The task’s color is initially black. After a session is recorded, the task’s color changes to orange or the color of the kanban column. Upon being moved to different kanban columns, the title will change color appropriately. The progress bar will change color depending on how close the sessions are to reaching the estimated time.  
Additionally, tasks in the Calendar component are color coded based on their kanban column location.

#### Acceptance

Pass.

### User Story \#3.1

As a user, I want to be able to login with social accounts (Google)

Acceptance Criteria:

1. ### Given the user clicks “log in with Google” and they have a google-linked account, they will be logged in as usual (manual)

2. ### Given the user clicks “register with Google” an account will be created with their Google name, email, and profile picture (manual)

3. ### Given the user initiates a calendar import flow but denies permission, the user will see an error indicating so (manual)

#### Test Case

This test case tests criteria 2

#### Input

1. Log out if logged in  
2. Go to “[https://ethandawes.github.io/Pleiades/registration](https://ethandawes.github.io/Pleiades/registration)”  
3. Click the “sign in with Google” button & follow the prompts  
4. Go to “[https://ethandawes.github.io/Pleiades/account](https://ethandawes.github.io/Pleiades/account)”

#### Output

- Observe that your profile picture matches Google  
- Observe that your name matches Google

#### Acceptance

Pass.

#### Test Case

This test case tests criteria 1

#### Input

1. Complete criteria 2 test case  
2. Log out if logged in  
3. Go to “[https://ethandawes.github.io/Pleiades/login](https://ethandawes.github.io/Pleiades/login)”  
4. Click the “sign in with Google” button & follow the prompts

#### Output

- Observe you are back in your account

#### Acceptance

Pass.

#### Test Case

This test case tests criteria 3

#### Input

1. Log in (any method)  
2. Open a project  
3. Go to the “Availability tab”  
4. Click “Import from google”  
5. Close the popup without doing anything

#### Output

- Observe an error toast saying “permission denied”

#### Acceptance

Pass.

### User Story \#5.1

As a user, I want to be able to upload my calendar from another program into Pleiades so that it can be used for finding team availabilities

Acceptance Criteria:

1. ### Given the user is on the availability tab and clicks “import from Google” or similar and their account is not connected to Google, they will be given an auth confirmation screen (manual)

2. ### Given the user is on the availability tab and clicks “import from Google” or similar and their account is connected to Google, their availability should be overwritten with “free” where there are no calendar events and “busy” where there is a calendar event (manual)

3. ### Given the user has calendar events marked as “free”, those should not be marked as “unavailable” in the availability calendar (manual)

#### Test Case

This test case tests all three acceptance criteria

#### Input

1. Add an event to your primary Google calendar  
2. Add another event to your primary Google calendar, mark it as “free”  
3. Open a project  
4. Go to the “availability” tab  
5. Click “Import from Google” button  
6. Follow the prompts, sign in with google

#### Output

- You should be prompted to grant the application additional permissions (the signin popup)  
- Your availability table should be overwritten with “available” (green) everywhere, except where there are events marked as “busy” (white)  
- Spots with events marked as “free” should be green

#### Acceptance

Pass.

### User Story \#2.3

As a user, I want to be able to invite people directly in the app using their email

Acceptance Criteria:

1. ### Given I enter an email to invite, that recipient will get an email with an invite code (manual)

2. ### Given I enter the name of a student at Purdue, a drop down appears with possible people matches. (manual)

3. ### Given I click on a suggested name, an email will be sent to their email retrieved from the Purdue directory. (manual)

#### Test Case

This test case tests all three acceptance criteria

#### Input

1. Log in, open project, open settings tab  
2. Click “Invite by email”  
3. Enter you email  
4. Click the chip with the email  
5. Type in your full name  
6. Click the chip with your name

#### Output

- Upon clicking a name chip, there will be a confirmation toast  
- You should receive two emails with an invite link  
- When entering your name, you should see a chip with your full name as it appears in the Purdue directory

#### Acceptance

Pass.

### User Story \#8.1

As a user, I want to be able to see how many people have created an account, visited the site, and hours logged to projects to provide credibility and provide feedback

Acceptance Criteria:

1. ### Given I visit the home page, I should see stats for how many people have created an account, visited the site, and hours logged (manual)

2. ### Given I perform one of the above actions, I should see the homepage updated to reflect the latest values on the next reload (manual)

3. ### Given the nav header is visible, I should see an additional button linking to a Google Form to provide feedback (manual)

#### Test Case

This test case tests all three acceptance criteria

#### Input

1. Visit the homepage and take note  
2. Reference previous testing reports on how to create an account and log hours  
3. Reload the site several times  
4. Visit the homepage again and take note  
5. Visit “[https://pleiadesapi.kimjammer.com/stats](https://pleiadesapi.kimjammer.com/stats)”

#### Output

- The stat counter should have increased by how many hours you worked (if less than an hour, visit the api endpoint to check  
- The stat counter should have increased by the number of new users you registered  
- The guest counter should have increased by how many times you loaded the page  
- Observe the new button in the nav bar that links to a feedback Google Form

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

