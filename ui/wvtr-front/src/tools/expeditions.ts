type ExpToGetFromBack = {
    category: string,
    key: string,
    imgURL: string,
    duration: number,
    costName: string,
    costNumber: number,
    canBeLaunched: boolean,
    order: number,
}

type ExpeditionCategory = {
    name: string,
    expeditions: ExpToGetFromBack[],
}

function buildExpeditionsCathegory(exps: ExpToGetFromBack[]): ExpeditionCategory[] {
    var rec: Record<string, ExpeditionCategory> = {}
    for (let i = 0; i < exps.length; i++) {
        let etgfb = rec[exps[i]!.category]
        if (etgfb != undefined) {
            etgfb.expeditions.push(exps[i]!)
        } else {
            let nExp = new Array<ExpToGetFromBack>
            nExp.push(exps[i]!)
            rec[exps[i]!.category] = {
                name: exps[i]!.category,
                expeditions: nExp,
            }
        }
    }

    var res = Object.entries(rec).map(([k, v]) => ({
        name: k,
        expeditions: v.expeditions
    }))

    res.sort((n1, n2) => {
        return n1.name.localeCompare(n2.name)
    });

    return res
}

export type {
    ExpeditionCategory,
    ExpToGetFromBack,
}

export {
    buildExpeditionsCathegory
}