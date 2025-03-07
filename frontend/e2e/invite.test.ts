import { test, expect } from "@playwright/test"

test("Create Project and open Project Page", async ({ page }) => {
    await page.goto("/home")
    await page.getByRole("link", { name: "Title Description" }).click()
    await new Promise(r => setTimeout(r, 1000))
    await expect(page.locator("h2")).toContainText("Title")
    await expect(page.getByRole("paragraph")).toContainText("Description")
})
