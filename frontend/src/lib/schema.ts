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

export type PollsResponse = {
    polls: string[]
}

export const taskformSchema = z.object({
    title: z.string().nonempty(),
    description: z.string(),
    due: z.string().date().optional(),
    timeEstimate: z.number().min(0).default(0),
    // this is an array of user ids NOT names. A custom component will make this easy for the user
    assignees: z.array(z.string()),
})

export type TaskFormSchema = typeof taskformSchema
export type TaskForm = z.infer<typeof taskformSchema>

export const pollformSchema = z.object({
    title: z.string().nonempty(),
    description: z.string().optional(),
    //options: z.string().nonempty(),
    dueDate: z.string().date(),
})

export type PollFormSchema = typeof pollformSchema
export type Poll = z.infer<typeof pollformSchema>
