import { IssueService } from "./issue.service";
import { TestBed } from "@angular/core/testing";

describe("IssueService", () => {
  let service: IssueService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(IssueService);
  });

  it("should be created", () => {
    expect(service).toBeTruthy();
  });
});
