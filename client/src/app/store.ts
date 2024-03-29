// store.ts
import { configureStore } from "@reduxjs/toolkit";
import userReducer from "./slices/user/user";
import personal_data from "./slices/profile/personal_data";
import tech_data from "./slices/profile/tech_data";

const store = configureStore({
  reducer: {
    user: userReducer,
    personal_data: personal_data,
    tech_data: tech_data,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export default store;
