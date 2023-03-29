describe("User that already exists", () => {
  it("Tries to login with user that already exists", () => {
    cy.visit("http://localhost:4200/register");
    cy.get("#mat-input-0").type("ramajanco");
    cy.get("#mat-input-1").type("Password123");
    cy.get("[id^=register-button]").click();
    cy.get("mat-card-footer").contains("Error: User already exists");
  });
});

describe("User with too short username", () => {
  it("Tries to register with username that is too short", () => {
    cy.visit("http://localhost:4200/register");
    cy.get("#mat-input-0").type("a");
    cy.get("#mat-input-1").type("Password123");
    cy.get("mat-error").contains("Username must be at least 4 characters long");
  });
});

describe("User with too long username", () => {
  it("Tries to register with username that is too long", () => {
    cy.visit("http://localhost:4200/register");
    cy.get("#mat-input-0").type("abcdefghijklmnopqrstuvwxyz");
    cy.get("#mat-input-1").type("Password123");
    cy.get("mat-error").contains("Username can't be longer then 20 characters");
  });
});

describe("No username", () => {
  it("Tries to register with no username", () => {
    cy.visit("http://localhost:4200/register");
    cy.get("#mat-input-0").clear();
    cy.get("#mat-input-1").type("Password123");
    cy.get("mat-error").contains("Username is required");
  });
});

describe("No password", () => {
  it("Tries to register with no password", () => {
    cy.visit("http://localhost:4200/register");
    cy.get("#mat-input-1").clear();
    cy.get("#mat-input-0").type("testpassword");
    cy.get("mat-error").contains("Password is required");
  });
});

describe("Password too short", () => {
  it("Tries to register with password that is too short", () => {
    cy.visit("http://localhost:4200/register");
    cy.get("#mat-input-1").type("abcdefg");
    cy.get("#mat-input-0").type("testuser");
    cy.get("mat-error").contains("Password must be at least 8 characters long");
  });
});

describe("Password too long", () => {
  it("Tries to register with password that is too long", () => {
    cy.visit("http://localhost:4200/register");
    cy.get("#mat-input-1").type("abcdefghijklmnopqrstuvwxyz");
    cy.get("#mat-input-0").type("testuser");
    cy.get("mat-error").contains("Password can't be longer then 25 characters");
  });
});

describe("Password with no uppercase", () => {
  it("Tries to register with password that has no uppercase", () => {
    cy.visit("http://localhost:4200/register");
    cy.get("#mat-input-1").type("password123");
    cy.get("#mat-input-0").type("testuser");
    cy.get("mat-error").contains("Must have lowercase, uppercase and a number");
  });
});

describe("Password with no lowercase", () => {
  it("Tries to register with password that has no lowercase", () => {
    cy.visit("http://localhost:4200/register");
    cy.get("#mat-input-1").type("PASSWORD123");
    cy.get("#mat-input-0").type("testuser");
    cy.get("mat-error").contains("Must have lowercase, uppercase and a number");
  });
});

describe("Password with no number", () => {
  it("Tries to register with password that has no number", () => {
    cy.visit("http://localhost:4200/register");
    cy.get("#mat-input-1").type("Password");
    cy.get("#mat-input-0").type("testuser");
    cy.get("mat-error").contains("Must have lowercase, uppercase and a number");
  });
});

describe("New User", () => {
  it("Registers new user and logs in", () => {
    cy.visit("http://localhost:4200/register");
    cy.get("#mat-input-0").type("testuser");
    cy.get("#mat-input-1").type("Password123");
    cy.get("[id^=register-button]").click();
    cy.url().should("eq", "http://localhost:4200/login");
    cy.get("[id^=username]").type("testuser");
    cy.get("[id^=password]").type("Password123");
    cy.get("[id^=login-button]").click();
    cy.url().should("eq", "http://localhost:4200/landing-page");
  });
});
