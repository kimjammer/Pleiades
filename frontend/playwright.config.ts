import { defineConfig, devices } from "@playwright/test"

export default defineConfig({
    webServer: {
        command: "npm run preview",
        port: 4173,
    },

    testDir: "e2e",

    projects: [
        {
            name: "setup",
            use: {
                ...devices["Desktop Chrome"],
                channel: "chromium",
            },
            testMatch: /.*\.setup\.ts/,
        },
        {
            name: "chromium",
            use: {
                ...devices["Desktop Chrome"],
                channel: "chromium",
                // Use prepared auth state.
                storageState: "test-results/.auth/user.json",
            },
            dependencies: ["setup"],
        },
    ],
})
