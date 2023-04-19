describe("Clicks on the map", () => {
    beforeEach(() => {
        cy.visit("http://localhost:4200/login");
        cy.get("[id^=username]").type("testuser1");
        cy.get("[id^=password]").type("Password123");
        cy.get("[id^=login-button]").click();
    })
    it("Clicks on the map and clicks submit", () => {
        cy.get("[id^=Gmap]").click("center");
        cy.get("[id^=submitButton]").click();
    });
});
describe("Signs out", () => {
    beforeEach(() => {
        cy.visit("http://localhost:4200/login");
        cy.get("[id^=username]").type("testuser1");
        cy.get("[id^=password]").type("Password123");
        cy.get("[id^=login-button]").click();
    })
    it("Clicks on the signout button", () => {
        cy.contains("Sign Out").click();
        cy.url().should("eq", "http://localhost:4200/home");
    });
});
describe("Test delete user", () => {
    it("Registers new user, logs in, and deletes user", () => {
      cy.visit("http://localhost:4200/register");
      cy.get("#mat-input-0").type("Test1");
      cy.get("#mat-input-1").type("Testpassword1");
      cy.get("[id^=register-button]").click();
      cy.url().should("eq", "http://localhost:4200/login");
      cy.get("[id^=username]").type("Test1");
      cy.get("[id^=password]").type("Testpassword1");
      cy.get("[id^=login-button]").click();
      cy.url().should("eq", "http://localhost:4200/landing-page");
      cy.get("[id^=delete]").click();
      cy.contains("Delete Account").click();
      cy.url().should("eq", "http://localhost:4200/home");
    });
    it("Tries to log in", () => {
        cy.visit("http://localhost:4200/login");
        cy.get("[id^=username]").type("Test1");
        cy.get("[id^=password]").type("Testpassword1");
        cy.get("[id^=login-button]").click();
        cy.get("div").contains("Error: No User Found");
    });
});
describe("Test the next button", () => {
    beforeEach(() => {
        cy.visit("http://localhost:4200/login");
        cy.get("[id^=username]").type("testuser1");
        cy.get("[id^=password]").type("Password123");
        cy.get("[id^=login-button]").click();
    })
    it("Clicks on the map, submits and clicks next", () => {
        cy.get("[id^=Gmap]").click("center");
        cy.get("[id^=submitButton]").click();
        cy.get("[id^=nextButton]").click();
        cy.url().should("eq", "http://localhost:4200/landing-page");
    });
});
