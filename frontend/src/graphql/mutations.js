import { gql } from "@apollo/client";

// Mutations
export const CREATE_RECIPE = gql`
  mutation CreateRecipe(
    $title: String!
    $description: String!
    $category: String!
    $prep_time: Int!
    $ingredients: [ingredients_insert_input!]!
    $steps: [steps_insert_input!]!
  ) {
    insert_recipes(
      objects: {
        title: $title
        description: $description
        category: $category
        prep_time: $prep_time
        ingredients: { data: $ingredients }
        steps: { data: $steps }
      }
    ) {
      returning {
        id
      }
    }
  }
`;

export const CREATE_USER = gql`
  mutation CreateUser($name: String!, $email: String!, $password: String!) {
    insert_users(objects: { name: $name, email: $email, password: $password }) {
      returning {
        id
      }
    }
  }
`;

export const CREATE_BOOKMARK = gql`
  mutation CreateBookmark($user_id: String!, $recipe_id: String!) {
    insert_bookmarks(objects: { user_id: $user_id, recipe_id: $recipe_id }) {
      returning {
        id
      }
    }
  }
`;

export const CREATE_CATEGORY = gql`
  mutation CreateCategory($name: String!) {
    insert_categories(objects: { name: $name }) {
      returning {
        id
      }
    }
  }
`;

export const CREATE_INGREDIENT = gql`
  mutation CreateIngredient($name: String!, $quantity: String!, $recipe_id: String!) {
    insert_ingredients(objects: { name: $name, quantity: $quantity, recipe_id: $recipe_id }) {
      returning {
        id
      }
    }
  }
`;

export const CREATE_STEP = gql`
  mutation CreateStep($step_number: Int!, $instruction: String!, $recipe_id: String!) {
    insert_steps(objects: { step_number: $step_number, instruction: $instruction, recipe_id: $recipe_id }) {
      returning {
        id
      }
    }
  }
`;

export const CREATE_IMAGE = gql`
  mutation CreateImage($url: String!, $recipe_id: String!) {
    insert_images(objects: { url: $url, recipe_id: $recipe_id }) {
      returning {
        id
      }
    }
  }
`;

export const CREATE_LIKE = gql`
  mutation CreateLike($user_id: String!, $recipe_id: String!) {
    insert_likes(objects: { user_id: $user_id, recipe_id: $recipe_id }) {
      returning {
        id
      }
    }
  }
`;

export const CREATE_RATING = gql`
  mutation CreateRating($user_id: String!, $recipe_id: String!, $rating: Int!) {
    insert_ratings(objects: { user_id: $user_id, recipe_id: $recipe_id, rating: $rating }) {
      returning {
        id
      }
    }
  }
`;

export const CREATE_COMMENT = gql`
  mutation CreateComment($user_id: String!, $recipe_id: String!, $content: String!) {
    insert_comments(objects: { user_id: $user_id, recipe_id: $recipe_id, content: $content }) {
      returning {
        id
      }
    }
  }
`;

export const CREATE_PURCHASE = gql`
  mutation CreatePurchase($user_id: String!, $recipe_id: String!, $amount: Float!) {
    insert_purchases(objects: { user_id: $user_id, recipe_id: $recipe_id, amount: $amount }) {
      returning {
        id
      }
    }
  }
`;

// Add more mutations for other tables as needed
