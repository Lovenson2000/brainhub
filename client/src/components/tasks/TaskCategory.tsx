import React from "react";

type TaskCategoryProps = {
  category: string;
};

export default function TaskCategory({ category }: TaskCategoryProps) {
  return (
    <div className="p-1.5 inline-block bg-indigo-500 bg-opacity-10 font-light rounded-sm text-sm text-slate-900">
      <h3>{category}</h3>
    </div>
  );
}
