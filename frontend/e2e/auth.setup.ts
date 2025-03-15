import { expect, test as setup } from "@playwright/test"
import "dotenv/config"

const authFile = "test-results/.auth/user.json"

setup("authenticate", async ({ page }) => {
    const url = "http://" + process.env.PUBLIC_API_HOST + "/resetDatabase"
    const res = await fetch(url, { mode: "cors", credentials: "include" })
    expect(res.status).toBe(200)

    await page.goto("http://localhost:4173/login")
    await page.getByPlaceholder("Email").fill("example@example.com")
    await page.getByPlaceholder("Password").fill("notthepassword")
    await page.getByRole("button", { name: "Login" }).click()
    await page.waitForURL("**/home")

    await page.context().storageState({ path: authFile })
})
