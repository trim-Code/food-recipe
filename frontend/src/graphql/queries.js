import { gql } from "@apollo/client";

// Queries
export const GET_RECIPES = gql`
  query GetRecipes {
    recipes {
      id
      title
      description
      created_at
    }
  }
`;

export const GET_USERS = gql`
  query GetUsers {
    users {
      id
      name
      email
    }
  }
`;

export const GET_BOOKMARKS = gql`
  query GetBookmarks {
    bookmarks {
      id
      user_id
      recipe_id
    }
  }
`;

export const GET_CATEGORIES = gql`
  query GetCategories {
    categories {
      id
      name
    }
  }
`;

export const GET_INGREDIENTS = gql`
  query GetIngredients {
    ingredients {
      id
      name
      quantity
      recipe_id
    }
  }
`;

export const GET_STEPS = gql`
  query GetSteps {
    steps {
      id
      step_number
      instruction
      recipe_id
    }
  }
`;

export const GET_IMAGES = gql`
  query GetImages {
    images {
      id
      url
      recipe_id
    }
  }
`;

export const GET_LIKES = gql`
  query GetLikes {
    likes {
      id
      user_id
      recipe_id
    }
  }
`;

export const GET_RATINGS = gql`
  query GetRatings {
    ratings {
      id
      user_id
      recipe_id
      rating
    }
  }
`;

export const GET_COMMENTS = gql`
  query GetComments {
    comments {
      id
      user_id
      recipe_id
      content
    }
  }
`;

export const GET_PURCHASES = gql`
  query GetPurchases {
    purchases {
      id
      user_id
      recipe_id
      amount
    }
  }
`;


