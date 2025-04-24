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

test("Create Task with all fields", async ({ page }) => {
    await resetDB()

    await page.goto(
        "http://localhost:4173/project?id=53ed4d28-9279-4b4e-9256-b1e693332625&tab=tasks",
    )
    await page.getByRole("button", { name: "Create a new task" }).click()
    await page.getByLabel("Title").fill("Write Speech")
    await page.getByLabel("Description").fill("Use the template")
    await page.getByLabel("Due date").fill("2025-05-01")
    await page.getByLabel("Time Estimate").fill("2")
    await page.getByRole("button", { name: "JS" }).click()
    await page.getByRole("button", { name: "Create!" }).click()
    await page.getByRole("button", { name: "Write Speech Use the template" }).click()
    await expect(page.getByText("Write Speech")).toBeVisible()
    await expect(page.getByText("Use the template")).toBeVisible()
    await expect(page.getByText("May 1")).toBeVisible()
    await expect(page.getByText("2 Hrs")).toBeVisible()
})

test("Create Task with NLP Time Estimate", async ({ page }) => {
    await resetDB()

    await page.goto(
        "http://localhost:4173/project?id=53ed4d28-9279-4b4e-9256-b1e693332625&tab=tasks",
    )
    await page.getByRole("button", { name: "Create a new task" }).click()
    await page.getByLabel("Title").fill("A Long Task 5 hours")
    await new Promise(r => setTimeout(r, 1000))
    await page.getByLabel("Description").fill("A long task description")
    await page.getByRole("button", { name: "Create!" }).click()
    await page.getByRole("button", { name: "A Long Task A long task description" }).click()
    await expect(page.getByText("5 Hrs")).toBeVisible()
})
