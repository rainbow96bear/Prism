import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import axios from "../../../configs/AxiosConfig";

interface User {
  user_id: string;
  nickname: string;
}

const initialState: User = {
  user_id: "",
  nickname: "",
};

// 비동기 작업을 처리하는 thunk 생성
export const fetchUser = createAsyncThunk<User>(
  "userInfo/fetchUser",
  async () => {
    const response = await axios.get<User>("/user/info", {
      withCredentials: true,
    });
    return response.data; // API 응답에서 필요한 데이터를 반환
  }
);

export const logout = createAsyncThunk("auth/logout", async () => {
  await axios.post("/OAuth/kakao/logout", {
    withCredentials: true,
  });

  // 로그아웃 후 리턴되는 데이터를 사용하려면 여기에 추가 코드 작성 가능
});

const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchUser.fulfilled, (state, action) => {
        // 비동기 작업이 성공하면 상태를 업데이트
        state.user_id = action.payload.user_id;
        state.nickname = action.payload.nickname;
      })
      .addCase(fetchUser.rejected, (state, action) => {
        state.user_id = "";
        state.nickname = "";
        console.error("Error fetching user info:", action.error.message);
      })
      .addCase(logout.fulfilled, (state) => {
        // 로그아웃이 성공하면 상태를 초기화
        state.user_id = "";
        state.nickname = "";
      })
      .addCase(logout.rejected, (state, action) => {
        console.error("Logout failed:", action.error.message);
      });
  },
});

export default userSlice.reducer;
