import { defineStore, acceptHMRUpdate } from "pinia";

interface ConstestRequest {
  name: string;
  slug: string;
  description: string;
}

interface ContestUpdateResponse {
  message: string;
  contest: ContestResponse;
}

export interface Contest {
  id: number | null;
  name: string;
  slug: string;
  description: string;
  problems: Problem[];
}

type ContestResponse = Contest;

/*
For later implementation
Contest -> {
  startTime: string | null;
  endTime: string | null;
  users: ContestUser[];
}

export interface ContestUser {
  id: number;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string | null;
  contest: Contest;
  contestId: number;
  user: User;
  userId: number;
  role: string;
  roleId: number;
}
*/

export const useContest = defineStore("contest", () => {
  // --- state ---
  const contests = ref<Contest[]>([]);

  // --- getters ---
  const getContestById = (id: number) => {
    return contests.value.find((contest) => contest.id === id);
  };

  // --- actions ---
  async function fetchContest(id: number) {
    const contest = getContestById(id);
    if (contest) {
      return contest;
    } else {
      try {
        const data = await useAPI<ContestResponse>(`/api/contests/${id}`);
        return data;
      } catch {
        console.error("Failed to fetch contest");
      }
    }
  }

  async function fetchAllContests() {
    try {
      const data = await useAPI<ContestResponse[]>(`/api/contests`);
      contests.value = data;
    } catch {
      console.error("Failed to fetch contests");
    }
  }

  async function createContest(request: ConstestRequest) {
    try {
      const data = await useAPI<ContestUpdateResponse>("/api/contests", {
        method: "POST",
        body: request,
      });

      await fetchAllContests();

      return data.contest.id;
    } catch (error) {
      throw useAPIError(error);
    }
  }

  async function updateContest(id: number, request: ConstestRequest) {
    try {
      const data = await useAPI<ContestUpdateResponse>(`/api/contests/${id}`, {
        method: "PUT",
        body: request,
      });

      await fetchAllContests();

      return data.contest.id;
    } catch (error) {
      throw useAPIError(error);
    }
  }

  return {
    contests,
    fetchContest,
    fetchAllContests,
    createContest,
    updateContest,
  };
});

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useContest, import.meta.hot));
}
