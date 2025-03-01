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
            testMatch: /.*\.setup\.ts/,
        },
        {
            name: "chromium",
            use: {
                ...devices["Desktop Chrome"],
                // Use prepared auth state.
                storageState: "test-results/.auth/user.json",
            },
            dependencies: ["setup"],
        },
    ],
})
