**Team 5 Project Charter**

Pleiades

**Team Members:** 

Cate, John, Henry, Ethan

**Problem Statement:**

Students at Purdue and similar universities are beset by the plight of bad group members, poor coordination, and bad time management. Teams would benefit strongly from tools that streamline teamwork. Pleiades is a group work “toolbox” that provides these tools such as finding open time slots in group members’ schedules, assigning roles and responsibilities, and tracking progress of group and individual tasks. Products already exist to achieve parts of these goals, but they only have one function or are aimed at individuals, making them unfit for group collaboration. Pleiades provides a unified platform specifically designed for short term group collaboration.

**Project Objectives:**



* Build a website with an intentional focus on a streamlined user experience that helps students to coordinate group work and communicate efficiently.
* Help people understand who’s responsible for what using a task assignment system.
* Develop a system where teams can input their availability and find time slots where everyone is available for meetups.
* Create a project and task tracking board that keeps team members informed on group progress.
    * Provide a burndown chart in the form of a stacked line plot by user.
    * Utilize a timer system for people to indicate how much time they’re spending on particular tasks.
* Create a voting system for making decisions on project direction using a highest mean rating system.
* Conduct user studies on friends and peers to ensure that the UX is as effective as possible by showing them a mockup and observing how they use it.

**Stakeholders:**



* Users (Purdue students in introductory CS courses/performing group work): They will want a user-friendly useful product.
* Developers (Cate, John, Henry, Ethan): we will want a pleasant developer experience.
* Product Owners (Cate, John, Henry, Ethan): we will want an impressive, useful product.
* Manager (Pratyush Das): will want the project to proceed on schedule and fulfill the requirements.

**Deliverables:**



* A Svelte based frontend with a design based on feedback from user studies.
* A Go-based backend for interacting with the database and serving the frontend.
* A postgresql database to store customer and business data.
* Utilize npm packages to build a service that decodes the ICAL format and generates team availability, with a page where users can input their availability manually.
