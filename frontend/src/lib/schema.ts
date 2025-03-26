import { z } from "zod"

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

export const formSchema = z.object({
    title: z.string().nonempty(),
    description: z.string(),
    due: z.date().optional(),
    estimate: z.number().min(0).optional(),
    // this is an array of user ids NOT names. A custom component will make this easy for the user
    asignees: z.array(z.string()),
})

export type FormSchema = typeof formSchema
export type Task = z.infer<typeof formSchema>
