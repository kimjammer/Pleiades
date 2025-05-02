**Team 5 Sprint 3 Retrospective**  
**Pleiades**  
Cate Harrison, John Kim, Henry Rovnyak, Ethan Dawes

**What went well?**

- We finished the application with all planned features and most of the stretch goals as well\!  
- Linking with external services, like email or google calendar, was successful and worked consistently.  
- We were able to work together to communicate changes in the CI/CD pipeline and hand off features for other team members to work with (example: emailing for forgotten passwords by John and then for invites by Ethan)  
- We found and fixed a merge/regression bug that didn’t get property merged (or reverted) where the tasks in the user’s calendar were referencing the wrong id before the sprint review by doing a practice run.

**What did not go well?**

- As our application size and complexity grew bigger, it became more difficult to find root causes of intermittent bugs.  
- Elaborating on above, I feel like there’s already a bit of technical debt. The `restEndpoints.go` file is 900 lines long, and there’s a bit of duplication within that.  
- CI/CD pipeline was annoying. It gave me anxiety about merge conflicts, made the history messy, we didn’t get to see other’s work and integrate it as quickly (ironic considering the name “continuous integration”)

**How can we improve?**

- Implement more thorough front-end testing that doesn’t require the backend.  
- Make time as a team to solicit user feedback along the design process to guide implementing features that are actually useful to users. It was on our radar, but since it wasn’t part of an acceptance criteria, it never got done.
