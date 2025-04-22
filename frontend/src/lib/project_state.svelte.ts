import { goto } from "$app/navigation"
import { base } from "$app/paths"
import { PUBLIC_API_HOST, PUBLIC_WS_PROTOCOL } from "$env/static/public"
import type { UserId } from "./schema"

class Mouse {
    x = $state(0)
    y = $state(0)
}

export let mouse = new Mouse()
document.addEventListener("mousemove", e => {
    mouse.x = e.pageX
    mouse.y = e.pageY
})
document.addEventListener("dragover", e => {
    mouse.x = e.pageX
    mouse.y = e.pageY
})

class ReactiveTesting {
    bruh = $state("X")
    list = $state(["A", "B", "C"])
    values: Record<string, string> = $state({
        x: "x",
        y: "y",
        z: "z",
    })
}

export async function connectToProject(projectId: string): Promise<ProjectState> {
    return new Promise((resolve, reject) => {
        let socket = new WebSocket(PUBLIC_WS_PROTOCOL + PUBLIC_API_HOST + "/ws")

        let state = new ProjectState(socket)
        let gotFirstProjectState = false
        let foundUserId = false

        let onmessage = function (this: WebSocket, message: MessageEvent) {
            let text = message.data

            if (typeof text != "string") {
                console.log("Server sent a non-string!", text, typeof text)
                return
            }

            console.log(text)

            if (text.startsWith("FAIL:")) {
                if (!gotFirstProjectState) {
                    reject(text)
                } else {
                    console.error(text)
                    state.error = text.slice(5)
                    state.showError = true
                }

                return
            }

            if (text.startsWith("WHOAMI: ")) {
                if (!gotFirstProjectState && !foundUserId) {
                    state.userId = text.slice(8)
                    foundUserId = true
                } else {
                    console.error(
                        "Server TXed the user ID twice or before sending the project state!",
                    )
                }

                return
            }

            if (text == "UNAUTHORIZED") {
                this.close()
                goto(base + "/login?project=" + projectId)
                return
            }

            if (text == "PROJECT ID DNE") {
                this.close()
                goto(base + "/home")
                return
            }

            if (!foundUserId) {
                console.error("Server did not TXed the user ID!")
                return
            }

            let serverState = JSON.parse(text)

            updateState(serverState, state)

            console.log($state.snapshot(state))

            if (!gotFirstProjectState) {
                gotFirstProjectState = true
                resolve(state)
            }
        }

        let onopen = () => {
            console.log("Connect")
            socket.send(projectId)
        }

        let joinSocket = (socket: WebSocket) => {
            socket.onmessage = onmessage
            socket.onopen = onopen
            socket.onclose = onclose
        }

        let onclose = () => {
            console.log("Close")
            setTimeout(() => {
                if (window.location.pathname != "/project") {
                    return
                }

                console.log("Trying to rejoin")
                socket = new WebSocket(PUBLIC_WS_PROTOCOL + PUBLIC_API_HOST + "/ws")
                state.socket = socket
                joinSocket(socket)
            }, 500)
        }

        joinSocket(socket)
    })
}

function updateState(serverResponse: any, state: ProjectState) {
    console.log(serverResponse)
    updateProject(serverResponse.Project, state)

    while (state.users.length > serverResponse.Users.length) {
        state.users.pop()
    }

    while (state.users.length < serverResponse.Users.length) {
        state.users.push(new UserInProject())
    }

    for (let i = 0; i < state.users.length; i++) {
        updateUser(serverResponse.Users[i], state.users[i])
    }
}

function updateUser(serverUser: any, user: UserInProject) {
    user.id = serverUser.Id
    user.leftProject = serverUser.LeftProject
    user.firstName = serverUser.FirstName
    user.lastName = serverUser.LastName
    user.email = serverUser.Email

    while (user.availability.length > serverUser.Availability.length) {
        user.availability.pop()
    }

    while (user.availability.length < serverUser.Availability.length) {
        user.availability.push(new Availability())
    }

    for (let i = 0; i < user.availability.length; i++) {
        updateAvailability(serverUser.Availability[i], user.availability[i])
    }
}

function updateAvailability(serverAvailability: any, availability: Availability) {
    availability.dayOfWeek = serverAvailability.DayOfWeek
    availability.startOffset = serverAvailability.StartOffset
    availability.endOffset = serverAvailability.EndOffset
}

function updateProject(serverProject: any, state: ProjectState) {
    state.id = serverProject.Id
    state.title = serverProject.Title
    state.description = serverProject.Description
    state.created = serverProject.Created
    state.demoButtonState = serverProject.DemoButtonState
    // TODO: Make the server send more comprehensive user info

    while (state.tasks.length > serverProject.Tasks.length) {
        state.tasks.pop()
    }

    while (state.tasks.length < serverProject.Tasks.length) {
        state.tasks.push(new Task())
    }

    for (let i = 0; i < state.tasks.length; i++) {
        updateTask(serverProject.Tasks[i], state.tasks[i])
    }

    while (state.polls.length > serverProject.Polls.length) {
        state.polls.pop()
    }

    while (state.polls.length < serverProject.Polls.length) {
        state.polls.push(new Poll())
    }

    for (let i = 0; i < state.polls.length; i++) {
        updatePoll(serverProject.Polls[i], state.polls[i])
    }
}

function updateTask(serverTask: any, task: Task) {
    task.id = serverTask.id
    task.title = serverTask.title
    task.description = serverTask.description
    task.dueDate = serverTask.dueDate
    task.kanbanColumn = serverTask.kanbanColumn
    task.timeEstimate = serverTask.timeEstimate
    task.completed = serverTask.completed

    while (task.sessions.length > (serverTask.sessions?.length ?? 0)) {
        task.sessions.pop()
    }

    while (task.sessions.length < (serverTask.sessions?.length ?? 0)) {
        task.sessions.push(new Session())
    }

    for (let i = 0; i < task.sessions.length; i++) {
        updateSession(serverTask.sessions[i], task.sessions[i])
    }

    task.assignees = serverTask.assignees
}

function updateSession(serverSession: any, session: Session) {
    session.id = serverSession.Id
    session.startTime = serverSession.StartTime
    session.endTime = serverSession.EndTime
    session.user = serverSession.User
}

function updatePoll(serverPoll: any, poll: Poll) {
    poll.id = serverPoll.Id
    poll.title = serverPoll.Title
    poll.description = serverPoll.Description
    poll.dueDate = serverPoll.DueDate

    while (poll.options.length > serverPoll.Options.length) {
        poll.options.pop()
    }

    while (poll.options.length < serverPoll.Options.length) {
        poll.options.push(new Option())
    }

    for (let i = 0; i < poll.options.length; i++) {
        updateOption(serverPoll.Options[i], poll.options[i])
    }
}

function updateOption(serverOption: any, option: Option) {
    option.id = serverOption.Id
    option.title = serverOption.Title
    option.likedUsers = serverOption.LikedUsers
    option.neutralUsers = serverOption.NeutralUsers
    option.dislikedUsers = serverOption.DislikedUsers
}

export class Availability {
    dayOfWeek = $state(0)
    startOffset = $state(0)
    endOffset = $state(0)
}

export class UserInProject {
    id: UserId = $state("")
    leftProject: boolean = $state(false)
    firstName: string = $state("")
    lastName: string = $state("")
    email: string = $state("")
    availability: Availability[] = $state([])
}

export class Session {
    id: string = $state("")
    startTime: number = $state(0)
    endTime: number = $state(0)
    user: UserId = $state("")
}

export class Task {
    id: string = $state("")
    title: string = $state("")
    description: string = $state("")
    dueDate: number = $state(0)
    kanbanColumn: string = $state("")
    timeEstimate: number = $state(0)
    completed: boolean = $state(false)
    sessions: Session[] = $state([])
    assignees: UserId[] = $state([])
}

export class Option {
    id: string = $state("")
    title: string = $state("")
    likedUsers: UserId[] = $state([])
    neutralUsers: UserId[] = $state([])
    dislikedUsers: UserId[] = $state([])
}

export class Poll {
    id: string = $state("")
    title: string = $state("")
    description: string = $state("")
    options: Option[] = $state([])
    dueDate: number = $state(0)
}

export class ProjectState {
    demoButtonState: "" | "a" | "b" = $state("")

    reactive_testing = $state(new ReactiveTesting())

    id: string = $state("")
    title: string = $state("")
    description: string = $state("")
    created: number = $state(0)
    users: UserInProject[] = $state([])
    tasks: Task[] = $state([])
    polls: Poll[] = $state([])
    userId: string = $state("")
    error: string = $state("")
    showError: boolean = $state(false)

    socket: WebSocket

    constructor(socket: WebSocket) {
        this.socket = socket
    }

    select(option: "" | "a" | "b") {
        this.socket.send(
            JSON.stringify({
                Name: "demoButtonState",
                Args: option,
            }),
        )
    }

    leave() {
        this.socket.send(
            JSON.stringify({
                Name: "leave",
                Args: {},
            }),
        )
        goto(base + "/home")
    }

    delete() {
        this.socket.send(
            JSON.stringify({
                Name: "delete",
                Args: {},
            }),
        )
        goto(base + "/home")
    }

    appendInProject<T>(key: string, value: T) {
        let message = JSON.stringify({
            Name: "append",
            Args: {
                Selector: key,
                NewValue: value,
            },
        })
        console.log(value, message)
        this.socket.send(message)
    }

    updateInProject<T>(key: string, value: T) {
        let message = JSON.stringify({
            Name: "update",
            Args: {
                Selector: key,
                NewValue: value,
            },
        })
        console.log(message)
        this.socket.send(message)
    }

    deleteInProject(key: string) {
        let message = JSON.stringify({
            Name: "remove",
            Args: {
                Selector: key,
            },
        })
        console.log(message)
        this.socket.send(message)
    }
}
