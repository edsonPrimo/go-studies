const [tasks, setTasks] = useState([{}])
const [taskInput, setTaskInput] = useState("")

handleForm() {
  const updatedTasks = [
    ...tasks,
    {
      nome: taskInput,
      id: Date.now() * Math.random(),
    }
  ]

  res = asopdksad()
  setTasks(res)
  localStorage.setItem("tasks", JSON.stringify(res))
}


const ok = [{ carlinhos }, { joao }]


[...ok, { }]