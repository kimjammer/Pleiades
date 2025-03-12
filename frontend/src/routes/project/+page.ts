import type { PageLoad } from "./$types"
import { superValidate } from "sveltekit-superforms"
import { formSchema } from "./tasks/schema"
import { zod } from "sveltekit-superforms/adapters"

export const load = (async () => {
    return {
        form: await superValidate(zod(formSchema)),
    }
}) satisfies PageLoad
