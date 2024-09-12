import { Navbar } from "../components/Navbar.js";
import "../styles/style.css";
import "../styles/post.css";
import "../styles/messenger.css";
import "../styles/chat.css";
import "../styles/users.css";
import { RenderPost } from "../components/Post";

import { GetAllUsers } from "../helpers/ServerRequests.js";
import { CONTAINER, ROOT, USERSCONTAINER } from "../index.js";
import { RenderPostFeed } from "../components/PostFeed.js";
import { RenderFilter } from "../components/Filter.js";
import { router } from "../router/Router.js";
import { OpenMessengers, Throttle } from "../components/Messenger.js";
import { LazyLoader } from "../helpers/LazyLoading.js";
import { UserList } from "../components/UserList.js";
import { CheckPosition } from "../helpers/LazyLoading.js";

export let POSTFEED;
export const loader = new LazyLoader()

window.addEventListener(
  "wheel",
  Throttle(async () => {
    console.log('jep testing throtthle')
    if(CheckPosition){
      fetchPosts(POSTFEED, await loader.Load())
    }
  }, 250)
);

window.addEventListener(
  "resize",
  Throttle(async () => {
    if(CheckPosition){
      fetchPosts(POSTFEED, await loader.Load())
    }
  }, 250)
);

export async function Home() {
  ROOT.innerHTML = "";
  CONTAINER.innerHTML = "";
  USERSCONTAINER.innerHTML = "";

  await Navbar();

  OpenMessengers.length = 0;
  POSTFEED = RenderPostFeed();
  ROOT.append(CONTAINER);

  //reset the lazyLoader
  loader.reset();
  fetchPosts(POSTFEED, await loader.Load());

  const usersData = await GetAllUsers();
  USERSCONTAINER.appendChild(UserList(usersData));
  CONTAINER.appendChild(USERSCONTAINER);

  const Filter = await RenderFilter();

  ROOT.appendChild(Filter);
  CONTAINER.appendChild(POSTFEED);
  CONTAINER.appendChild(USERSCONTAINER);

  const selectBlock = document.querySelector(".select-block");
  if (selectBlock) {
    selectBlock.addEventListener("click", displayCategoryModal);
    const dropdown = document.querySelector(".select-dropdown");
    const selectArrow = document.querySelector(".select-arrow");

    function displayCategoryModal() {
      dropdown.classList.toggle("showSelectModal");
      selectArrow.classList.toggle("select-arrow-rotate");
    }
  }
}

export async function fetchPosts(PostFeed, data) {
  try {
    const postsData = data

    if (postsData) {
      postsData.forEach((post) => {
        const postLink = document.createElement("div");
        postLink.classList.add("post-link");

        postLink.addEventListener("click", () => {
          history.pushState({}, "", `post/${post.post.id}`);
          router();
        });

        if (post.categories) {
          postLink.appendChild(RenderPost(post.post, post.categories));
        } else {
          postLink.appendChild(RenderPost(post.post));
        }

        PostFeed.appendChild(postLink);
      });
    } else {
      // Handle case when response is not OK
      console.log("Error: Response not OK");
    }
  } catch (error) {
    // Handle errors that occurred during the fetch
    console.error("Error during fetch:", error);
  }
}

///

//create a link and calls function to create post out of provided data
export function ProcessPostData(post) {
  const postLink = document.createElement("div");
  postLink.addEventListener("click", () => {
    history.pushState({}, "", `post/${post.post.id}`);
    router();
  });

  postLink.classList.add("post-link");

  if (post.categories) {
    postLink.appendChild(RenderPost(post.post, post.nickname));
  } else {
    postLink.appendChild(RenderPost(post.post));
  }

  return postLink;
}