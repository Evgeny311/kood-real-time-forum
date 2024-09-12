import { GLOBAL_URL } from "../config";

export class LazyLoader{
  constructor(){
    this.timestamp;
    this.categoryId;
    this.index;
    this.reset();
  }

  setIndex(value){
    this.index = value;
  }

  getIndex(){
    return this.index;
  }

  setTimestamp(value){
    this.timestamp = value
  }

  setCategory(value){
    this.categoryId = value;
  }

  getCategory() {
    return this.categoryId
  }

  reset(){
    this.index = 0;
    this.timestamp = new Date().toISOString();
    this.categoryId = "";
  }

  async Load(){
      const requestData = {
        current_index: this.index,
        page_open_time_stamp: this.timestamp,
        category_id: this.categoryId,
      };

      try {
        const response = await fetch(GLOBAL_URL + "/api/v1/jwt/posts", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
            body: JSON.stringify(requestData),
        });

        if (response.status === 401) {
            console.error("Unauthorized: User needs to authenticate");
            return null; // or throw an error if appropriate
        }

        const data = await response.json();
        this.index++;
        return data.data;
    } catch (error) {
        console.error("While lazy loading an error occurred! ", error);
        return null; // or throw an error if appropriate
    }
    }
}


export function Throttle(callee, timeout) {
  let timer = null;

  return function perform(...args) {
    if (timer) return;
    timer = setTimeout(() => {
      callee(...args);

      clearTimeout(timer);
      timer = null;
    }, timeout);
  };
}

export function CheckPosition() {
  const height = document.body.offsetHeight;
  const screenHeight = window.innerHeight;

  const scrolled = window.scrollY;
  const threshold = height - screenHeight / 4;

  const position = scrolled + screenHeight;

  if (position >= threshold) {
    return true
  }
  return false;
}
