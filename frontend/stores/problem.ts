interface ProblemRequest {
  name: string;
  slug: string;
  author: string;
  description: string;
}

interface ProblemUpdateResponse {
  message: string;
  problem: ProblemResponse;
}

export interface Problem {
  id: number | null;
  name: string;
  slug: string;
  author: string;
  admin: User;
  description: string;
}

type ProblemResponse = Problem;

/* 10 DEFAULT PROBLEMS */
const defaultProblems: Problem[] = [
    {
        id: 1,
        name: "Two Sum",
        slug: "two-sum",
        author: "Arvid Kristoffersson",
        admin: {
            id: 1,
            username: "ArvidKristoffersson",
            avatarUrl: "https://picsum.photos/60/60",
            permissions: [],
            email: ""
        },
        description: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.",
    },
    {
        id: 2,
        name: "Reverse String",
        slug: "reverse-string",
        author: "David Björklund",
        admin: {
            id: 2,
            username: "davidbjorklund",
            avatarUrl: "https://picsum.photos/60/60",
            permissions: [],
            email: ""
        },
        description: "Write a function that reverses a string. The input string is given as an array of characters.",
    },
    {
        id: 3,
        name: "Palindrome Number",
        slug: "palindrome-number",
        author: "Valter Mann",
        admin: {
            id: 3,
            username: "Quois",
            avatarUrl: "https://picsum.photos/60/60",
            permissions: [],
            email: ""
        },
        description: "Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.",
    },
    {
        id: 4,
        name: "Merge Sorted Array",
        slug: "merge-sorted-array",
        author: "LiamoOo",
        admin: {
            id: 4,
            username: "liamthorell",
            avatarUrl: "https://picsum.photos/60/60",
            permissions: [],
            email: ""
        },
        description: "Given two sorted integer arrays, merge them into a single sorted array.",
    },
];

export const useProblem = defineStore("problem", () => {
  // --- state ---
  const problems = ref<Problem[]>([]);

  // --- getters ---
  const getProblemById = (id: number) => {
    return problems.value.find((problem) => problem.id === id);
  };
  
  const getProblemBySlug = (slug: string) => {
    return problems.value.find((problem) => problem.slug === slug);
  };

  // --- actions ---
  async function fetchProblem(slug: string) {
    const problem = getProblemBySlug(slug);
    if (problem) {
      return problem;
    } else {
        /* TEMPORARY REPLACEMENT */
        const defaultProblem = defaultProblems.find((problem) => problem.slug === slug);
        if (defaultProblem) {
            return defaultProblem;
        }
        console.error("Problem not found in default problems");
        return null;
        /* END TEMPORARY REPLACEMENT */
        /*
      try {
        const data = await useAPI<ProblemResponse>(`/api/problems/${slug}`);
        return data;
      } catch {
        console.error("Failed to fetch problem");
      }
        */
    }
  }

  async function fetchAllProblems() {
    /* TEMPORARY REPLACEMENT */
    if(problems.value.length < 4) {
        problems.value = defaultProblems;
        return;
    }
    /* END TEMPORARY REPLACEMENT */
    /*
    try {
      const data = await useAPI<ProblemResponse[]>(`/api/problems`);
      problems.value = data;
    } catch {
      console.error("Failed to fetch problems");
    }
    */
  }

  async function createProblem(request: ProblemRequest) {
    /* TEMPORARY REPLACEMENT */
    problems.value.push({
        id: problems.value.length + 1,
        name: request.name,
        slug: request.slug,
        author: request.author,
        admin: {
            id: 5,
            username: "Your Username",
            email: "",
            avatarUrl: "https://picsum.photos/60/60",
            permissions: [],
        },
        description: request.description,
    });
    
    return request.slug;
    /* END TEMPORARY REPLACEMENT */
    /*
    try {
      const data = await useAPI<ProblemUpdateResponse>("/api/problems", {
        method: "POST",
        body: request,
      });

      await fetchAllProblems();

      return data.problem.id;
    } catch (error) {
      throw useAPIError(error);
    }
      */
  }

  async function updateProblem(slug: string, request: ProblemRequest) {
    /* TEMPORARY REPLACEMENT */
    const index = problems.value.findIndex((problem) => problem.slug === slug);
    if (index !== -1) {
        problems.value[index].name = request.name;
        problems.value[index].slug = request.slug;
        problems.value[index].author = request.author;
        problems.value[index].description = request.description;
    }
    /* END TEMPORARY REPLACEMENT */

    /*
    try {
      const data = await useAPI<ProblemUpdateResponse>(`/api/problems/${id}`, {
        method: "PUT",
        body: request,
      });

      await fetchAllProblems();

      return data.problem.id;
    } catch (error) {
      throw useAPIError(error);
    }
      */
  }

  return {
    problems,
    fetchProblem,
    fetchAllProblems,
    createProblem,
    updateProblem,
  };
});

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useProblem, import.meta.hot));
}
