import { formSchema } from "$lib/schema"
import { superValidate } from "sveltekit-superforms"
import { zod } from "sveltekit-superforms/adapters"
import type { PageLoad } from "./$types"

export const load = (async () => {
    return {
        form: await superValidate(zod(formSchema)),
    }
}) satisfies PageLoad
