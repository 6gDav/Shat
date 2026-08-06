const isClient = typeof window !== 'undefined';

let initialName = "";
if (isClient) {
  initialName = sessionStorage.getItem("userName") || "";
}

let nameState = $state<string>(initialName);

export const user = {
  get userName() {
    return nameState;
  },
  set userName(value: string) {
    nameState = value;
    if (isClient) {
      sessionStorage.setItem("userName", value);
    }
  }
};