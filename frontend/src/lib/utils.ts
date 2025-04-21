import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs))
}

// Adapted from https://gist.github.com/ca0v/73a31f57b397606c9813472f7493a940
export function debounce<T extends Function>(cb: T, wait = 20) {
    let h: any
    let callable = (...args: any) => {
        clearTimeout(h)
        h = setTimeout(() => cb(...args), wait)
    }
    return <T>(<any>callable)
}

export const GOOGLE_OAUTH_CLIENT_ID = "104619113313-21gij0dmhi857nk0c4tf7gjbgfc6jg82"
