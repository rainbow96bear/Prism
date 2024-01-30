import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import axios from "../../../configs/AxiosConfig";

interface User {
  id: string;
  nickname: string;
  oneLineIntroduce: string;
  hashtag: string[];
}

const initialState: User = {
  id: "",
  nickname: "",
  oneLineIntroduce: "",
  hashtag: [],
};

// 비동기 작업을 처리하는 thunk 생성
export const getProfile = createAsyncThunk<User>(
  "userInfo/getProfile",
  async (id) => {
    const response = await axios.get<User>(
      `/users/profiles/personaldatas/${id}`,
      {
        withCredentials: true,
      }
    );
    return response.data; // API 응답에서 필요한 데이터를 반환
  }
);

export const logout = createAsyncThunk("auth/logout", async () => {
  await axios.post(
    "/oauth/kakao/logout",
    {},
    {
      withCredentials: true,
    }
  );

  // 로그아웃 후 리턴되는 데이터를 사용하려면 여기에 추가 코드 작성 가능
});

const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(getProfile.fulfilled, (state, action) => {
        // 비동기 작업이 성공하면 상태를 업데이트
        state.id = action.payload.id;
        state.nickname = action.payload.nickname;
        state.oneLineIntroduce = action.payload.oneLineIntroduce;
      })
      .addCase(getProfile.rejected, (state, action) => {
        state.id = "";
        state.nickname = "";
        state.oneLineIntroduce = "";
        console.error("Error fetching user info:", action.error.message);
      })
      .addCase(logout.fulfilled, (state) => {
        // 로그아웃이 성공하면 상태를 초기화
        state.id = "";
        state.nickname = "";
      })
      .addCase(logout.rejected, (state, action) => {
        console.error("Logout failed:", action.error.message);
      });
  },
});

export default userSlice.reducer;
