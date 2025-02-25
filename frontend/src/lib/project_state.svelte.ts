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

export async function connectToProject(key: string, projectId: string): Promise<ProjectState> {
    return new Promise((resolve, reject) => {
        let socket = new WebSocket("ws://" + PUBLIC_API_HOST + "/ws")

        let state = new ProjectState()
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

            if (text == "INVALID TOKEN") {
                window.location.replace("/login")
            }

            if (text == "PROJECT ID DNE") {
                window.location.replace("/projects")
            }

            let new_state = JSON.parse(text)

            console.log(new_state)

            // TODO: Update the old state

            if (!gotFirstProjectState) {
                gotFirstProjectState = true
                resolve(state)
            }
        }

        socket.onopen = () => {
            socket.send(key)
            socket.send(projectId)
        }

        socket.onclose = () => {
            console.log("Closed!")
        }
    })
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
    button_state: "enabled" | "a" | "b" = $state("enabled")

    reactive_testing = $state(new ReactiveTesting())

    id: string = $state("")
    title: string = $state("")
    description: string = $state("")
    users: UserInProject[] = $state([])
    tasks: Task[] = $state([])
    polls: Poll[] = $state([])

    select(option: "a" | "b") {
        if (this.button_state == "enabled") {
            this.button_state = option
        }
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
