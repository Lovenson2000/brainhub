"use client";
import React, { useState } from "react";

export default function UserDetails() {
  const [user] = useState(() => {
    if (typeof window !== "undefined") {
      const storedUser = localStorage.getItem("user");
      return storedUser ? JSON.parse(storedUser) : null;
    }
    return null;
  });
  return (
    <div className="p-4">
      {user ? (
        <h1 className="text-2xl font-bold">
          Welcome {user.firstname} {user.lastname}
        </h1>
      ) : (
        <h1 className="text-2xl font-bold">No user found</h1>
      )}
    </div>
  );
}
