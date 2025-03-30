import { taskformSchema } from "$lib/schema"
import {pollformSchema} from "$lib/schema";
import { superValidate } from "sveltekit-superforms"
import { zod } from "sveltekit-superforms/adapters"
import type { PageLoad } from "./$types"

export const load = (async () => {
    return {
        taskform: await superValidate(zod(taskformSchema)),
        pollform: await superValidate(zod(pollformSchema))
    }
}) satisfies PageLoad
