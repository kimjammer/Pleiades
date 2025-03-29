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
    estimate: z.number().min(0).optional(),
    // this is an array of user ids NOT names. A custom component will make this easy for the user
    asignees: z.array(z.string()),
})

export type FormSchema = typeof taskformSchema
export type Task = z.infer<typeof taskformSchema>

export const pollformSchema = z.object({
    title: z.string().nonempty(),
    description: z.string().optional(),
    dueDate: z.string().date(),
    options: z.string().nonempty()
})

export type PollFormSchema = typeof pollformSchema
export type Poll = z.infer<typeof pollformSchema>;
