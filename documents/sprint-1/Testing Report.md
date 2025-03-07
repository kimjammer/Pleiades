**Team 5 Sprint 1 Testing Report**  
**Pleiades**  
Cate Harrison, John Kim, Henry Rovnyak, Ethan Dawes

Setup:

1. Cd frontend  
2. Npm i  
3. Npm run dev  
4. (in another session) cd backend  
5. go run \-tags TEST .  
6. Open localhost:5173 in browser


### User Story \#2.8

As a user, I want to be able to see a list of groups I am in so that I can choose which group I want to view

Acceptance Criteria:

1. Given that the UI has been created, when the user navigates to the user home page, they should see a grid of projects that they are a part of. (Automated)  
2. Given that the UI has been created, when the user navigates to the user home page, they should see a button to create a new project. (Automated)  
3. Given that the user is part of at least 1 project, when the user clicks on a project’s card, it should navigate them to that project’s home page. (Automated)  
4. Given that the user is not logged in, when the user navigates to the user home page, they should be redirected to the login screen. (**Manual**)

#### Test Case

This test case tests acceptance criteria 4, and whether or not the user is properly routed to the login screen if they are not logged in while trying to access the user home page.

#### Input

Open the website, do not log in, and directly access /home.

#### Output

The user is redirected to the /login page.

#### Acceptance

Pass.

### User Story \#2.10

As a user, I want to create and view a group information panel with a project name, project description, and links in the description automatically converted to anchor tags

Acceptance Criteria:

1. Given that the project exists, when the user navigates to the project page, they see the project name. (Automated)  
2. Given that the project exists, when the user navigates to the project page, they see the project description. (Automated)  
3. Given that the project description contains links, when the user navigates to the project page, they see the links in the description as anchor tags. (**Manual**)  
4. Given that the user is not logged in, when the user navigates to the project page, they should be redirected to the login screen. (**Manual**)

#### Test Case

This test case tests acceptance criteria 3, confirming that sections of the description that appear to be URLs are converted into anchor tags

#### Input

Open the website, login, go to the user home page, and create a new project with the description “My epic new project, info at https://www.google.com”. Open that project’s dashboard, and click on the URL.

#### Output

The URL in the description is a clickable anchor tag and clicking it opens a new tab with that page.

#### Acceptance

Pass.

#### Test Case

This test case tests acceptance criteria 4, confirming that if a logged out user attempts to access a project page, they are redirected to the login screen.

#### Input

Open the website, login, go to the user home page, and open (or create and open) a project page. Copy the project page URL. Go to the landing screen and log out. Paste the URL into the nav bar and attempt to access the project page

#### Output

The user is redirected to the login page.

#### Acceptance

Pass.

### User Story \#1.1

As a user, I want to be able to create an account with an email, real name, and password.

Acceptance Criteria:

1. Given that an email is already registered, it cannot be registered again for a different account and an error message will pop up.  
2. Given that the user does not have an account, they can create an account with an unregistered email.  
3. Given that the user has successfully registered, they should be directed to their home screen.  
4. Given that an email is already registered, users can navigate back to the log-in screen.  
5. Given that a user enters a password without capital letters, lowercase letters, special characters, and below a specific length, it will not be accepted during registration.

#### Test Case

This test case tests acceptance criteria 5, given that a user enters a password without capital letters, lowercase letters, special characters, and below a specific length, it will not be accepted during registration.

#### Input

Open the website, and on the landing page select “Create an Account.” Enter “Password43” as the password. This demonstrates the requirement for special characters.  
Next, enter “Password$” This demonstrates the requirement for numbers.  
Next, enter “Pass$43” This demonstrates the requirement for \>= 8 character passwords.  
Next, enter “Password$43” This demonstrates an acceptable password.

#### Output

The registration page should show the user an error message for the first 3 passwords entered, as they are not acceptable. The last password should not display the password requirements error message

#### Acceptance

Pass.

#### Test Case

This test case tests acceptance criteria 2, given that the user does not have an account, they can create an account with an unregistered email, acceptance criteria 3, given that the user has successfully registered, they should be directed to their home screen.

#### Input

NOTE: This test assumes that the email “example@example.com” IS NOT registered. If this is not the case, enter a different unregistered email.  
Open the website, and on the landing page select “Create an Account.” Enter example@example.com in the email box, some first and last name, and “Password.43” as the password.

#### Output

The website should direct you to the home screen. Additionally, the website should have a cookie with a token “token”, which indicates that the user is logged in.

#### Acceptance

Pass.

#### Test Case

This test case tests acceptance criteria 1, given that an email is already registered, it cannot be registered again for a different account and an error message will pop up.

#### Input

NOTE: This test assumes that the email “example@example.com” IS registered. If this is not the case, enter a different registered email.  
Open the website, and on the landing page select “Create an Account.” Enter example@example.com in the email box, some first and last name, and “Password.43” as the password.

#### Output

An error should appear indicating that the email is already registered.

#### Acceptance

Pass.

#### Test Case

This test case tests acceptance criteria 4, given that an email is already registered, users can navigate back to the log-in screen.

#### Input

Open the website, and on the landing page select “Create an Account.” Click “Login”

Output  
The website should direct the user to the login page

#### Acceptance

Pass.

### User Story \#1.2

As a user, I want to be able to log in with an email and password so that I can see information for me and my team.

Acceptance Criteria:

1. Given that an email is registered, the account can be logged in to.  
2. Given that the user is successfully logged in, they should be directed to their home screen.  
3. Given that the user submits invalid login data, an error message will pop up.

#### Test Case

This test case tests acceptance criteria 1, given that an email is registered, the account can be logged in to, and acceptance criteria 2, given that the user is successfully logged in, they should be directed to their home screen.

Input  
NOTE: This test case assumes the email “example@example.com” IS REGISTERED as per the user story 1.1 tests. If this is not the case, please register an email before running this.  
Open the website, and click “Login”. Enter “example@example.com” and “Password.43” as the email and password. Click “Login”

#### Output

The user should be directed to their home page.

#### Acceptance

Pass.

#### Test Case

This test case tests acceptance criteria 3, given that the user submits invalid login data, an error message will pop up.

Input  
Open the website and click “Login”. Enter “bad@email.com” and “password” as the email and password, and click “Login”.

#### Output

An error message pops up informing the user that the login information is invalid.

#### Acceptance

Pass.

### User Story \#2.6

As a user, I want to be able to leave a group.

Acceptance Criteria:

1. Upon clicking the “leave” button, a confirmation modal will appear and clicking “Confirm” will cause the user to leave the group and redirect to the homepage. The group will not appear there. (Manual)  
2. Upon clicking the “leave” button, a confirmation modal will appear and clicking “Cancel” will cause the modal to close and nothing to happen. (Manual)  
3. An invitation link will allow the user to rejoin the group (Manual)

#### Test Case

This test case tests all three acceptance criteria.

#### Input

Assume that the user is logged in. On the user home page, create a new project and enter the project. Click on the \`settings\` tab and click \`Leave\`, then click \`Cancel\`. Click \`Invite\` and copy the invite link. Click \`Leave\`, and click \`Confirm\`. Visit the copied invite link and click \`Accept\`. Return to the user home page.

#### Output

Upon clicking \`cancel\`, the modal will close. Upon clicking \`confirm\`, the user will be redirected to the home page and the project will not appear. Upon clicking \`Accept\`, the user will be redirected to the project page. Upon returning to the home page, the project will appear in the list of projects.

#### Acceptance

Pass.

### User Story \#2.7

As a user, I want to be able to delete a group.

Acceptance Criteria:

1. Upon clicking the “delete” button, a confirmation modal will appear and clicking “Confirm” will cause the group to be deleted and the user redirected to the main page. The group will not appear there. (Manual)  
2. Upon clicking the “delete” button, a confirmation modal will appear and clicking “Cancel” will cause the modal to close and nothing to happen. (Manual)  
3. Invitation links will not work after deleting the group. (Manual)

#### Test Case

This test case tests all three acceptance criteria.

#### Input

Assume that the user is logged in. On the user home page, create a new project and enter the project. Click on the \`settings\` tab and click \`Delete\`, then click \`Cancel\`. Click \`Invite\` and copy the invite link. Click \`Delete\`, and click \`Confirm\`. Visit the copied invite link.

#### Output

Upon clicking \`cancel\`, the modal will close. Upon clicking \`confirm\`, the user will be redirected to the home page and the project will not appear. Upon visiting the invite link, the invite link will display that the link is invalid.

#### Acceptance

Pass.

### User Story \#2.2

As a user I want to generate an invite link to the group so that I can collaborate with other students

Acceptance Criteria:

1. Given the user is logged in, when the ‘invite’ route is visited, it will return a token expiring in a week that can be used to invite new members (automated)  
2. Given the user is logged out, when the ‘/invite’ route is visited, it will return a 401 unauthorized error (automated)  
3. Given the user is logged in, when they click the share button, the application should contact the server, get the token, and display a URL and QR code to join (manual)

#### Test Case

This tests acceptance criteria 3, that the QR code and link appear

#### Input

Open the website, log in, open a project, change to the “settings” tab, and click the “invite” button.

#### Output

Use a mobile device to confirm that the QR code is the same as the one shown on screen. Also confirm there is a link in the format “/invite?id=\<UUID\>”

### User Story \#2.4

As a user I want to click on invite links and automatically be added to that group

Acceptance Criteria:

1. Given an expired invite link, when the user visits it, they will see a notice that it’s expired (manual)  
2. Using a link while signed out will redirect to login page (user story 2.5 will improve the flow) (manual)  
3. Given a valid invite link, when the user visits it, they will see a confirmation page with basic project info and an option to accept or decline (manual)  
4. Given an invalid invite link, when the user visits it, they will see a message indicating the link is invalid (manual)

#### Test Case

Acceptance criteria 3, 1

#### Input

Open the website, log in, open a project, change to the “settings” tab, and click the “invite” button. Copy the link. Log out and log into a different account (complete under 20 seconds before it expires. In production, it expires after a week). Paste the link, click the accept button.

#### Output

Observe that there is the project name, description, and two buttons to accept and decline. Click the accept and observe how you are directed to the newly added project page

#### Acceptance

Pass.

#### Test Case

Acceptance criteria 4

#### Input

Open the website, log in, and visit http://localhost:5173/join?id=dne

#### Output

It should say “Invalid or expired invite”

#### Acceptance

Pass.

#### Test Case

Acceptance criteria 1

#### Input

Open the website, log in, open a project, change to the “settings” tab, and click the “invite” button. Click the link. Wait 25 seconds (Testing mode only. In production, it expires after a week). Reload the page

#### Output

It should say “Invalid or expired invite”

#### Acceptance

Pass.

### User Story \#5.2

As a user I want to be able to manually set availabilities if I don’t use a traditional calendar

Acceptance Criteria:

1. Given the UI is implemented, when the user navigates to the availability page, they should see a table. On the X axis are dates, the Y axis is times in 15-minute increments, and the cells can be checked and unchecked to indicate whether or not they are available during that time block.  
2. Given the UI is implemented, when a user navigates to the availability page, they should see a drop-down to change their time zone  
3. Given time zones work correctly, when the user switches their timezone, the times and dates should shift to reflect their current time zone

#### Test Case

Acceptance criteria 1

#### Input

Log in, click a project, go to the availability tab. Click on a single cell on the lefthand table. Click and drag a rectangle within the table. Click again on a green cell.

#### Output

Observe the existence of the availability table and how it conforms to the specification. Observe that when clicking or dragging, the cells turn green. Observe that clicking a green cell turns white again.

#### Acceptance

Pass.

#### Test Case

Acceptance criteria 2 & 3

#### Input

Log in, click a project, go to the availability tab. Click the timezone dropdown, initially set to your current timezone (likely “America/New\_York”). Choose a different timezone

#### Output

Observe the existence of the dropdowns & confirm that **every** timezone is present: [https://en.wikipedia.org/wiki/List\_of\_tz\_database\_time\_zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). After clicking, observe the y-axis labels on the availability input change to reflect the current timezone

#### Acceptance

Pass.

### User Story \#5.3

As a user I want to be able to view everyone’s availability as a calendar in a weekly view

Acceptance Criteria:

1. Given the group calendar works, when multiple different people input their availability, the combined availability calendar should be darker in areas where the most people are available (manual)  
2. Given the group calendar works, when hovering over a time block, the user should see a list of members who are available or not during that time block (manual)  
3. Given the group calendar works, as you or other members input their availability in real-time, changes should be reflected immediately (manual)

#### Test Case

This test case tests all 3 acceptance criteria.

#### Input

Open the website, log in, open a project, change to the “availability” tab. In a separate incognito window, sign in to a different account, join the same project as the first user (see user story 2.2 for how to do that), and also switch to the availability tab. On one account, click and drag a region on the lefthand availability input. On the other account, click and drag an overlapping region on the lefthand availability input.

#### Output

Observe that where the availability from each account overlaps, the cell is darker. Observe that when hovering over colored cells in the righthand grid, there is a tooltip with who is available. Observe that as you edit the availability, you see that on the other user in real time.

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