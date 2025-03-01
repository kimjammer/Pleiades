import { test, expect } from "@playwright/test"

test("Create Project Modal appears", async ({ page }) => {
    await page.goto("/home")
    await page.getByRole("button", { name: "Create a new project" }).click()
    await expect(page.getByRole("dialog", { name: "Create new Project" })).toBeVisible()
})

test("Create Project Modal stays on error", async ({ page }) => {
    await page.goto("/home")
    await page.getByRole("button", { name: "Create a new project" }).click()
    await page.getByRole("button", { name: "Create!" }).click()
    await expect(page.getByRole("dialog", { name: "Create new Project" })).toBeVisible()
})

test("Create Project Modal dismisses on success", async ({ page }) => {
    await page.goto("/home")
    await page.getByRole("button", { name: "Create a new project" }).click()
    await page.getByRole("textbox", { name: "Title" }).fill("Title")
    await page.getByRole("textbox", { name: "Description" }).fill("Description")
    await page.getByRole("button", { name: "Create!" }).click()
    await expect(page.getByRole("dialog", { name: "Create new Project" })).toHaveCount(0)
})
