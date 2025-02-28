import { PUBLIC_API_HOST } from "$env/static/public"

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
        let socket = new WebSocket("ws://" + PUBLIC_API_HOST + "/ws")

        let state = new ProjectState(socket)
        let gotFirstProjectState = false

        socket.onmessage = message => {
            let text = message.data

            if (typeof text != "string") {
                console.log("Server sent a non-string!", text, typeof text)
                return
            }

            if (text.startsWith("FAIL:")) {
                if (!gotFirstProjectState) {
                    reject(text)
                } else {
                    console.error(text)
                }

                return
            }

            if (text == "PROJECT ID DNE") {
                window.location.replace("/home")
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

        socket.onopen = () => {
            socket.send(projectId)
        }

        socket.onclose = () => {
            console.log("Closed!")
        }

        socket.onerror = e => {
            // Is there a way to tell if the error is caused by "unauthorized"?
            window.location.replace("/login")
        }
    })
}

function updateState(serverResponse: any, state: ProjectState) {
    state.id = serverResponse.Id
    state.title = serverResponse.Title
    state.description = serverResponse.Description
    state.demoButtonState = serverResponse.DemoButtonState
    // TODO: Make the server send more comprehensive user info

    while (state.tasks.length > serverResponse.Tasks.length) {
        state.tasks.pop()
    }

    while (state.tasks.length < serverResponse.Tasks.length) {
        state.tasks.push(new Task())
    }

    for (let i = 0; i < state.tasks.length; i++) {
        updateTask(serverResponse.Tasks[i], state.tasks[i])
    }

    while (state.polls.length > serverResponse.Polls.length) {
        state.polls.pop()
    }

    while (state.polls.length < serverResponse.Polls.length) {
        state.polls.push(new Poll())
    }

    for (let i = 0; i < state.polls.length; i++) {
        updatePoll(serverResponse.Polls[i], state.polls[i])
    }
}

function updateTask(serverTask: any, task: Task) {
    task.id = serverTask.Id
    task.title = serverTask.Title
    task.description = serverTask.Description
    task.dueDate = serverTask.DueDate
    task.kanbanColumn = serverTask.KanbanColumn
    task.timeEstimate = serverTask.TimeEstimate
    task.completed = serverTask.completed

    while (task.sessions.length > serverTask.Sessions.length) {
        task.sessions.pop()
    }

    while (task.sessions.length < serverTask.Sessions.length) {
        task.sessions.push(new Session())
    }

    for (let i = 0; i < task.sessions.length; i++) {
        updateSession(serverTask.Sessions[i], task.sessions[i])
    }

    task.assignees = serverTask.Assignees
}

function updateSession(serverSession: any, session: Session) {
    session.id = serverSession.Id
    session.startTime = serverSession.StartTime
    session.startTime = serverSession.EndTime
    session.user = serverSession.User
}

function updatePoll(serverPoll: any, poll: Poll) {
    poll.id = serverPoll.Id
    poll.title = serverPoll.Title
    poll.description = serverPoll.Description

    while (poll.options.length > serverPoll.Options.length) {
        poll.options.pop()
    }

    while (poll.options.length < serverPoll.Options.length) {
        poll.options.push(new Option())
    }

    for (let i = 0; i < poll.options.length; i++) {
        updateOption(serverPoll.Sessions[i], poll.options[i])
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
    id: string = $state("")
    firstName: string = $state("")
    lastName: string = $state("")
    availability: Availability[] = $state([])
}

export class Session {
    id: string = $state("")
    startTime: string = $state("") // TODO: Change to correct type
    endTime: string = $state("") // TODO: Change to correct type
    user: string = $state("")
}

export class Task {
    id: string = $state("")
    title: string = $state("")
    description: string = $state("")
    dueDate: string = $state("") // TODO: Change to correct type
    kanbanColumn: string = $state("")
    timeEstimate: number = $state(0)
    completed: boolean = $state(false)
    sessions: Session[] = $state([])
    assignees: string[] = $state([])
}

export class Option {
    id: string = $state("")
    title: string = $state("")
    likedUsers: string[] = $state([])
    neutralUsers: string[] = $state([])
    dislikedUsers: string[] = $state([])
}

export class Poll {
    id: string = $state("")
    title: string = $state("")
    description: string = $state("")
    options: Option[] = $state([])
}

export class ProjectState {
    demoButtonState: "" | "a" | "b" = $state("")

    reactive_testing = $state(new ReactiveTesting())

    id: string = $state("")
    title: string = $state("")
    description: string = $state("")
    users: UserInProject[] = $state([])
    tasks: Task[] = $state([])
    polls: Poll[] = $state([])

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

    appendInProject(key: string, value: any) {
        traverseObject(
            this,
            toPath(key),
            (object, key) => {
                if (object[key] == undefined) {
                    object[key] = value
                } else if (Array.isArray(object[key])) {
                    object[key].push(value)
                } else {
                    throw `Unsure how to append ${value} to ${key}. Did you mean 'updateInProject'?`
                }
            },
            (array, idx) => {
                if (Array.isArray(array[idx])) {
                    array[idx].push(value)
                } else {
                    throw `Unsure how to append ${value} to ${key}. Did you mean 'updateInProject'?`
                }
            },
        )
    }

    updateInProject(key: string, value: any) {
        traverseObject(
            this,
            toPath(key),
            (object, key) => {
                object[key] = value
            },
            (array, idx) => {
                array[idx] = value
            },
        )
    }

    deleteInProject(key: string) {
        console.log(key, toPath(key))
        traverseObject(
            this,
            toPath(key),
            (object, key) => {
                delete object[key]
            },
            (array, idx) => {
                array.splice(idx, 1)
            },
        )
    }
}

function toPath(data: string): string[] {
    // The `replaceAll` replaces instances of `things[5].etc` with `things.5.etc`
    return data.replaceAll(/\[([^\]]*)\]/g, (_, idx) => `.${idx}`).split(".")
}

function traverseObject(
    object: any,
    key: string[],
    toAffectObject: (final_spot: any, final_key: string) => void,
    toAffectArray: (final_spot: Array<any>, final_key: number) => void,
) {
    if (Array.isArray(object)) {
        if (key.length == 1) {
            toAffectArray(object, parseInt(key[0]))
        } else {
            traverseObject(object[parseInt(key[0])], key.slice(1), toAffectObject, toAffectArray)
        }
    } else {
        if (key.length == 1) {
            toAffectObject(object, key[0])
        } else {
            traverseObject(object[key[0]], key.slice(1), toAffectObject, toAffectArray)
        }
    }
}
