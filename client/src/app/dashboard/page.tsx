"use client";
import React, { useState } from "react";

export default function Page() {
  const [user, setUser] = useState(() => {
    const storedUser =
      typeof window !== "undefined" ? localStorage.getItem("user") : null;
    return storedUser ? JSON.parse(storedUser) : null;
  });

  return (
    <div className="p-4">
      {user ? (
        <h1 className="text-2xl font-bold">
          Welcome {user.firstname} {user.lastname}
        </h1>
      ) : (
        <h1 className="text-2xl font-bold">Loading...</h1>
      )}
    </div>
  );
}
