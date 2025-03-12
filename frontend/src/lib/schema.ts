export type ProjectsResponse = {
    projects: minimalProject[]
}

export type minimalProject = {
    id: string
    title: string
    description: string
}

export type newProjectRequest = {
    title: string
    description: string
}
