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
    // TODO: Don't hardcode this and read it from the URL
    // We can't do that now because the ports are different
    let socket = new WebSocket("ws://localhost:8080/ws")

    socket.onmessage = message => {}

    socket.onopen = () => {
        socket.send(key)
        socket.send(projectId)
    }

    socket.onclose = () => {
        console.log("Closed!")
    }

    return new Promise((resolve, reject) => {})
}

export class ProjectState {
    button_state: "enabled" | "a" | "b" = $state("enabled")

    reactive_testing = $state(new ReactiveTesting())

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
