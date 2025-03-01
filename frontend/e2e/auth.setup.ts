import { test as setup } from "@playwright/test"
import "dotenv/config"

const authFile = "test-results/.auth/user.json"

setup("authenticate", async ({ page }) => {
    await page.goto("/")
    await page.getByRole("button", { name: "Login as some user" }).click()
    await page.context().storageState({ path: authFile })
})
