import { goto } from "$app/navigation"
import { PUBLIC_PROTOCOL, PUBLIC_API_HOST } from "$env/static/public"
import { base } from "$app/paths"

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

export async function tryJoinProject() {
    const params = new URLSearchParams(location.search)
    const projectId = params.get("project")
    if (projectId === null) return false
    await joinProject(projectId)
    return true
}
