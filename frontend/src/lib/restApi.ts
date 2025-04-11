import { goto } from "$app/navigation"
import { base } from "$app/paths"
import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"

export async function joinProject(projectId: string) {
    const resp = await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/join" + location.search, {
        mode: "cors",
        credentials: "include",
    })
    if (resp.status === 200) {
        goto(base + "/project?id=" + projectId)
    }

    return resp
}

/**
 * @returns true if success, false otherwise. Used to determine if should continue to home
 */
export async function tryJoinProject() {
    const params = new URLSearchParams(location.search)
    const projectId = params.get("project")
    const joinId = params.get("id")
    if (projectId && !joinId) {
        goto(base + "/project?id=" + projectId)
        return true
    }
    if (projectId === null) return false
    await joinProject(projectId)
    return true
}

/** Log an event and return a promise resolving to boolean success */
export function recordEvent(name: string, count = 1) {
    return fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + `/event?name=${name}&value=${count}`, {
        method: "GET",
        mode: "cors",
    }).then(res => res.status === 200)
}
