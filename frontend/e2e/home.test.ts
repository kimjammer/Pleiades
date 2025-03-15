import { test, expect } from "@playwright/test"

async function resetDB() {
    const url = "http://" + process.env.PUBLIC_API_HOST + "/resetDatabase"
    const res = await fetch(url, { mode: "cors", credentials: "include" })
    expect(res.status).toBe(200)
}

test("Create Project Modal appears", async ({ page }) => {
    await resetDB()

    await page.goto("/home")
    await page.getByRole("button", { name: "Create a new project" }).click()
    await expect(page.getByRole("dialog", { name: "Create new Project" })).toBeVisible()
})

test("Create Project Modal stays on error", async ({ page }) => {
    await resetDB()

    await page.goto("/home")
    await page.getByRole("button", { name: "Create a new project" }).click()
    await page.getByRole("button", { name: "Create!" }).click()
    await expect(page.getByRole("dialog", { name: "Create new Project" })).toBeVisible()
})

test("Create Project Modal dismisses on success", async ({ page }) => {
    await resetDB()

    await page.goto("/home")
    await page.getByRole("button", { name: "Create a new project" }).click()
    await page.getByRole("textbox", { name: "Title" }).fill("Title")
    await page.getByRole("textbox", { name: "Description" }).fill("Description")
    await page.getByRole("button", { name: "Create!" }).click()
    await expect(page.getByRole("dialog", { name: "Create new Project" })).toHaveCount(0)
})

test("Create Project and open Project Page", async ({ page }) => {
    await resetDB()

    await page.goto("/home")
    await page.getByRole("link", { name: "Test Project Test Description" }).click()
    await new Promise(r => setTimeout(r, 1000))
    await expect(page.locator("h2")).toContainText("Test Project")
    await expect(page.getByRole("paragraph")).toContainText("Test Description")
})
